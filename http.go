package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 这是结构体
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// 这是切片
var albums = []album {
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	fmt.Println(newAlbum)
	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbums(c *gin.Context) {
	//c.String(http.StatusOK, "hello")
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func mul(c *gin.Context) {

	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")

	fmt.Printf("ids: %v; names: %v", ids, names)
}

func loginEndpoint(c *gin.Context)  {
	c.String(http.StatusOK,"login")
}

func submitEndpoint(c *gin.Context)  {
	c.String(http.StatusOK, "submit")
}

func readEndpoint(c *gin.Context)  {
	c.String(http.StatusOK, "read")
}

func main() {
	router := gin.Default() //默认路由 含有中间件 Logger Recovery

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.POST("/mul_post", mul)

	// 简单组： v1
	v1 := router.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.GET("/read", readEndpoint)
	}

	router.Run("localhost:8080")
}