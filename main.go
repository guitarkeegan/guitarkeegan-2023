package main

import (
	"fmt"
	"log"
	"net/http"

	blog "guitarkeegan-2023.com/v2/model"

	"github.com/gin-gonic/gin"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	router := gin.Default()
	router.Use(gin.ErrorLogger())
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.Static("/videos", "./videos")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main", nil)
	})

	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.POST("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})

	router.GET("/projects", func(c *gin.Context) {
		c.HTML(http.StatusOK, "projects.html", gin.H{
			"title": "PocketPR",
			"p1":    "What is the project?",
			"p2":    "What was my role in the project?",
			"p3":    "Where is the project at now?",
		})
	})

	router.GET("/blog", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blog.html", nil)
	})

	blog.CreateDBConnection()

	res, err := blog.GetPosts(3)

	checkErr(err)
	fmt.Println(res)

	router.Run(":8080")
}
