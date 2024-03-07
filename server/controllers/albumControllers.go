package controllers

import (
	"encoding/json"
	"fmt"
	"io"
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

func GetAlbums (c *gin.Context) {
	fmt.Println("Total albums: ======================");

	c.JSON(http.StatusOK, gin.H{
		"albums": albums,
	})

}

func GetAlbum (c *gin.Context) {
	fmt.Println("++++++++++++++++++++");
	id := c.Query("id")

	for _, album := range albums {
		if(album.ID == id) {
			c.JSON(http.StatusOK, gin.H{
				"album": album,
			})
		}
	}

	c.JSON(http.StatusBadRequest, gin.H {
		"album": "Album not found",
	})
}

func PostAlbum(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Request": "Bad request",
		})
		return
	}


	var newAlbum album

	err = json.Unmarshal(jsonData, &newAlbum)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Request": "Invalid JSON data",
		})
		return
	}

	albums = append(albums, newAlbum)
}

func DeleteAlbum (c *gin.Context) {
	fmt.Print("Deleting =======================");

}