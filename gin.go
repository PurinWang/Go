package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getHtml(c *gin.Context){
	c.HTML(http.StatusOK, "test.html",gin.H{
		"title":"Hey",
	})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.LoadHTMLGlob("view/*")
	router.GET("/html",getHtml)

	router.LoadHTMLGlob("templates/**/*")
	router.GET("/html2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/user.html", gin.H{
			"title": "Users",
		})
	})
	router.Run("localhost:8080")
}
