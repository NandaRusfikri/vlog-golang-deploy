package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

var Barang = []Product{
	{ID: 1, Name: "Raket Yonek KW", Description: "John Coltrane", Price: 56.99},
	{ID: 2, Name: "Tas Eager Ultraligth", Description: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Name: "Xiaomi Redmi note 4 pro", Description: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/", index)
	router.GET("/product", getAlbums)
	router.GET("/product/:id", getAlbumByID)
	router.POST("/product", postAlbums)

	router.Run("localhost:80")
}

// Barang slice to seed record Barang data.

func index(c *gin.Context) {
	operatingSystem := runtime.GOOS
	switch operatingSystem {
	case "windows":
		fmt.Println("Windows")
	case "darwin":
		fmt.Println("MAC operating system")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", operatingSystem)
	}
	hostname, _ := os.Hostname()
	jsonData := map[string]string{
		"title":            "Welcome to my Golang Web API",
		"message":          "Gin web framework is awesome",
		"hostname":         hostname,
		"Operating System": operatingSystem,
	}

	c.JSON(http.StatusOK, jsonData)
}

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, Barang)
}

func postAlbums(c *gin.Context) {
	var newAlbum Product

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	Barang = append(Barang, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	intVar, _ := strconv.Atoi(id)

	for _, a := range Barang {
		if a.ID == intVar {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
