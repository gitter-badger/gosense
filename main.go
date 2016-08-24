package main

import (
	"database/sql"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang/groupcache/lru"
	"github.com/naoina/toml"
	_ "github.com/netroby/mysql"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type ShowMessage interface {
	ShowMessage(c *gin.Context)
}
type msg struct {
	msg string
}
type umsg struct {
	msg string
	url string
}

type VBlogItem struct {
	aid            int
	title          sql.NullString
	content        sql.NullString
	publish_time   sql.NullString
	publish_status sql.NullInt64
	views          int
}

/*
* ShowMessage with template
 */
func (m *msg) ShowMessage(c *gin.Context) {
	c.HTML(http.StatusOK, "message.html", gin.H{
		"message": template.HTML(m.msg),
	})
}

func (m *umsg) ShowMessage(c *gin.Context) {

	c.HTML(http.StatusOK, "message.html", gin.H{
		"message": template.HTML(m.msg),
		"url":     m.url,
	})
}

func GetMinutes() string {
	return time.Now().Format("200601021504")
}

func GetDB(config *appConfig) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=30s&charset=utf8mb4",
		config.Db_user, config.Db_password, config.Db_host, config.Db_port, config.Db_name))
	if err != nil {
		panic(err.Error())
	}
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(100)
	return db
}

type appConfig struct {
	Db_host          string
	Db_port          int
	Db_name          string
	Db_user          string
	Db_password      string
	Admin_user       string
	Admin_password   string
	Site_name        string
	Site_description string
	ObjectStorage    struct {
		Aws_access_key_id     string
		Aws_secret_access_key string
		Aws_region            string
		Aws_bucket            string
		Cdn_url               string
	}
}

func GetConfig() *appConfig {
	f, err := os.Open("config.toml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var config appConfig
	if err := toml.Unmarshal(buf, &config); err != nil {
		panic(err)
	}
	return &config
}

var (
	Config    *appConfig
	DB        *sql.DB
	Cache     *lru.Cache
	CacheSize int = 8192
)

func main() {

	Config = GetConfig()
	DB = GetDB(Config)
	Cache = lru.New(CacheSize)

	r := gin.Default()
	r.StaticFS("/assets", assetFS())
	store := sessions.NewCookieStore([]byte("gssecret"))
	r.Use(sessions.Sessions("mysession", store))
	tplsname := []string{
		"index.html",
		"about.html",
		"add-blog.html",
		"admin.list.blog.html",
		"admin-files.html",
		"admin-login.html",
		"edit-blog.html",
		"message.html",
		"search.html",
		"view.html",
		"donate.html",
	}
	var t *template.Template
	for i := 0; i < len(tplsname); i++ {
		tn := fmt.Sprintf("templates/%s", tplsname[i])
		tpl, err := Asset(tn)
		if err == nil {
			var tmpl *template.Template
			if t == nil {
				t = template.New(tplsname[i])
			}
			if tplsname[i] == t.Name() {
				tmpl = t
			} else {
				tmpl = t.New(tplsname[i])
			}
			_, _ = tmpl.Parse(string(tpl))
		}
	}
	r.SetHTMLTemplate(template.Must(t, nil))

	fc := new(FrontController)
	r.GET("/", fc.HomeCtr)
	r.GET("/about", fc.AboutCtr)
	r.GET("/view/:id", fc.ViewCtr)
	r.GET("/view.php", fc.ViewAltCtr)
	r.GET("/ping", fc.PingCtr)
	r.GET("/search", fc.SearchCtr)
	r.GET("/countview/:id", fc.CountViewCtr)

	ac := new(AdminController)
	admin := r.Group("/admin")
	{
		admin.GET("/", ac.ListBlogCtr)
		admin.GET("/login", ac.LoginCtr)
		admin.POST("/login-process", ac.LoginProcessCtr)
		admin.GET("/logout", ac.LogoutCtr)
		admin.GET("/addblog", ac.AddBlogCtr)
		admin.POST("/save-blog-add", ac.SaveBlogAddCtr)
		admin.GET("/listblog", ac.ListBlogCtr)
		admin.GET("/deleteblog/:id", ac.DeleteBlogCtr)
		admin.POST("/save-blog-edit", ac.SaveBlogEditCtr)
		admin.GET("/editblog/:id", ac.EditBlogCtr)
		admin.GET("/files", ac.Files)
		admin.POST("/fileupload", ac.FileUpload)
	}

	a := new(api)
	api := r.Group("/api")
	{
		api.GET("/", a.index)
		api.GET("view/:id", a.view)
	}
	rss := new(RSS)
	r.GET("/rss.php", rss.Alter)
	r.GET("/rss", rss.Out)
	endless.ListenAndServe(":8080", r)
}
