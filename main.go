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
		c.HTML(http.StatusOK, "about.html", nil)
	})

	router.GET("/projects", func(c *gin.Context) {
		res, err := blog.GetProjects(3)

		checkErr(err)

		c.HTML(http.StatusOK, "projects.html", res)
	})

	router.GET("/blog", func(c *gin.Context) {

		res, err := blog.GetPosts(3)

		checkErr(err)

		c.HTML(http.StatusOK, "blog.html", res)
	})

	router.GET("/blog/:id", func(c *gin.Context) {
		id := c.Param("id")

		post, err := blog.GetPostById(id)

		fmt.Println(post)

		checkErr(err)

		if post.Title == "" {
			c.HTML(http.StatusBadRequest, "post.html", gin.H{
				"Title":       "none found",
				"Description": "-",
				"ID":          "#",
			})
			return
		}

		c.HTML(http.StatusOK, "post.html", post)
	})

	blog.CreateDBConnection()

	router.Run(":8080")
}
