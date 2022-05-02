package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var categories = []category{
	{ID: "1", Name: "Bar"},
}

func getCategories(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, categories)
}

func postCategory(c *gin.Context) {
	var newCategory category

	if err := c.BindJSON(&newCategory); err != nil {
		return
	}

	categories = append(categories, newCategory)
	c.IndentedJSON(http.StatusCreated, newCategory)
}

func getCategoryByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range categories {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "category not found"})
}

func main() {
	router := gin.Default()
	router.GET("/api/category", getCategories)
	router.POST("/api/category", postCategory)
	router.GET("/api/category/:id", getCategoryByID)

	router.Run("localhost:8080")
}
