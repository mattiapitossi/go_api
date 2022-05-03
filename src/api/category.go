package api

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

func GetCategories(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, categories)
}

func PostCategory(c *gin.Context) {
	var newCategory category

	if err := c.BindJSON(&newCategory); err != nil {
		return
	}

	categories = append(categories, newCategory)
	c.IndentedJSON(http.StatusCreated, newCategory)
}

func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range categories {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "category not found"})
}
