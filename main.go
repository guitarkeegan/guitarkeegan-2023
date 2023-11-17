package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.ErrorLogger())
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.Static("/videos", "./videos")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.POST("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", nil)
	})

	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.POST("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})

	router.Run(":8080")
}
