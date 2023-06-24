package main

import (
	"products-api/api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	productsRouting(router)
	categoriesRouting(router)

	router.Run()
}

func productsRouting(router *gin.Engine) {
	router.GET("/products/:productID", controller.GetProductByID)
	router.GET("/products/category/:categoryID", controller.GetProductsByCategoryID)
	router.POST("/products", controller.AddProduct)
	router.PUT("/products", controller.UpdateProduct)
	router.DELETE("/products/:productID", controller.DeleteProduct)
}

func categoriesRouting(router *gin.Engine) {
	router.GET("/categories", controller.GetCategories)
	router.GET("/categories/:categoryID", controller.GetCategoryByID)
	router.POST("/categories", controller.AddCategory)
	router.PUT("/categories", controller.UpdateCategory)
	router.DELETE("/categories/:categoryID", controller.DeleteCategory)
}
