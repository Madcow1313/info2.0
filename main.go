package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var data Data
	data.fields = make(map[string]any)
	data.fillMain()
	router := gin.Default()
	loadFiles(router)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", data.fields)
	})
	router.GET("/about.html", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "about.html", data.fields)
	})
	router.GET("/data.html", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "data.html", data.fields)
	})
	router.GET("/data.html/:create", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "about.html", data.fields)
	})
	router.Run()
}
