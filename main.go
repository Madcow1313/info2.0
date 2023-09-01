package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("index.html")
	router.StaticFile("/logo.svg", "logo.svg")
	router.StaticFile("/static/js/main.js", "./static/js/main.js")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"logo": "logo.svg",
		})
	})
	router.Run()
}
