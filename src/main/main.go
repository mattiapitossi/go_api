package main

import (
	"github.com/gin-gonic/gin"
	"go_web/src/api"
)

func main() {
	router := gin.Default()

	categories := router.Group("/api/category")
	{
		categories.GET("", api.GetCategories)
		categories.POST("", api.PostCategory)
		categories.GET("/:id", api.GetCategoryByID)
	}

	router.Run("localhost:8080")
}
