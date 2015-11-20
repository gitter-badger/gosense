package main

import (
	"github.com/gin-gonic/gin"	
	"net/http"
)
// RSS controll group
type RSS struct {
	
} 
// Out Render and output RSS
func (rss *RSS) Out(c *gin.Context) {
	c.HTML(http.StatusOK, "rss.html", gin.H{})
}