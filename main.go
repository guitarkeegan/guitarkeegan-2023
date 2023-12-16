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

// trigger koyeb
func main() {

	gin.SetMode(gin.ReleaseMode)
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
		if c.GetHeader("HX-Request") == "true" {
			c.HTML(http.StatusOK, "about.html", nil)
			return
		}

		c.HTML(http.StatusOK, "about-dir.html", nil)
	})

	router.GET("/projects", func(c *gin.Context) {
		res, err := blog.GetProjects(3)

		checkErr(err)

		if c.GetHeader("HX-Request") == "true" {
			c.HTML(http.StatusOK, "projects.html", res)
			return
		}
		c.HTML(http.StatusOK, "projects-dir.html", res)
	})

	router.GET("/projects/:id", func(c *gin.Context) {
		id := c.Param("id")

		project, err := blog.GetProjectById(id)

		fmt.Println(project)

		checkErr(err)

		if project.Title == "" {
			c.HTML(http.StatusBadRequest, "projects.html", gin.H{
				"Title":       "none found",
				"Description": "-",
				"ID":          "#",
			})
			return
		}
		c.HTML(http.StatusOK, "project-single.html", project)
	})

	router.GET("/blog", func(c *gin.Context) {

		res, err := blog.GetPosts(3)

		checkErr(err)

		if c.GetHeader("HX-Request") == "true" {
			c.HTML(http.StatusOK, "blog.html", res)
			return
		}
		c.HTML(http.StatusOK, "blog-dir.html", res)
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

	router.Run("0.0.0.0:8080")
}
