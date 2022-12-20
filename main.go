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

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	index := -1

	for i, a := range albums {
		if a.ID == id {
			index = i
			break
		}
	}
	if index != -1 {
		albums = append(albums[:index], albums[index+1:]...)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "deleted successfully"})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cannot find id"})
}

func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	index := -1
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	for i, a := range albums {
		if a.ID == id {
			index = i
			break
		}
	}
	if index != -1 {
		albums[index] = newAlbum
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Updated successfully"})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.DELETE(("/albums/:id"), deleteAlbum)
	router.PUT("albums/:id", updateAlbum)

	router.Run("localhost:8080")
}
