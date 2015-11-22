package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

// RSS controll group
type RSS struct {
}

// Out Render and output RSS
func (rss *RSS) Out(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		fmt.Println(err)
	}
	page -= 1
	if page < 0 {
		page = 0
	}

	prev_page := page
	if prev_page < 1 {
		prev_page = 1
	}

	rpp := 20
	offset := page * rpp
	CKey := fmt.Sprintf("rss-home-page-%d-rpp-%d", page, rpp)
	var blogListSlice []apiBlogList
	val, ok := Cache.Get(CKey)
	if val != nil && ok == true {
		fmt.Println("Ok, we found cache, Cache Len: ", Cache.Len())
		blogListSlice = val.([]apiBlogList)
	} else {
		rows, err := DB.Query("Select aid, title from top_article where publish_status = 1 order by aid desc limit ? offset ? ", &rpp, &offset)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		var (
			aid   sql.NullString
			title sql.NullString
		)
		blogListSlice = make([]apiBlogList, 0) //Must be zero slice
		var aBlog apiBlogList
		for rows.Next() {
			err := rows.Scan(&aid, &title)
			if err != nil {
				fmt.Println(err)
			}
			aBlog.Aid = aid.String
			aBlog.Title = title.String
			blogListSlice = append(blogListSlice, aBlog)
		}
		err = rows.Err()
		if err != nil {
			fmt.Println(err)
		}
		Cache.Add(CKey, blogListSlice)
	}
	c.HTML(http.StatusOK, "rss.html", gin.H{
		"bloglist": blogListSlice,
		"title":    "title",
	})
}
