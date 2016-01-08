package main

import (
	"database/sql"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang/groupcache/lru"
	"html/template"
)

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
