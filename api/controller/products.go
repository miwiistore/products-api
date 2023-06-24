package controller

import (
	"products-api/api/domain"
	"products-api/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProductByID(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		c.IndentedJSON(400, "Invalid parameter.")
		return
	}

	product, err := service.GetProductByID(int64(productID))
	if err != nil {
		c.IndentedJSON(500, err.Error())
		return
	}

	c.IndentedJSON(200, product)
}

func GetProductsByCategoryID(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		c.IndentedJSON(400, "Invalid parameter.")
		return
	}

	products, err := service.GetProductsByCategoryID(int16(categoryID))
	if err != nil {
		c.IndentedJSON(500, err.Error())
		return
	}

	c.IndentedJSON(200, products)
}

func AddProduct(c *gin.Context) {
	var newProduct domain.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.IndentedJSON(400, "Input body is invalid.")
		return
	}

	newProduct, err := service.AddProduct(newProduct)
	if err != nil {
		c.IndentedJSON(500, err.Error())
		return
	}

	c.IndentedJSON(200, newProduct)
}

func UpdateProduct(c *gin.Context) {
	var currentProduct domain.Product
	if err := c.ShouldBindJSON(&currentProduct); err != nil {
		c.IndentedJSON(400, "Input body is invalid.")
		return
	}

	if currentProduct.ProductID == 0 {
		c.IndentedJSON(400, "Missing product_id. Field is required.")
		return
	}

	currentProduct, err := service.UpdateProduct(currentProduct)
	if err != nil {
		c.IndentedJSON(500, err.Error())
		return
	}

	c.IndentedJSON(200, currentProduct)
}

func DeleteProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		c.IndentedJSON(400, "Invalid parameter.")
		return
	}

	err = service.DeleteProduct(int64(productID))
	if err != nil {
		c.IndentedJSON(500, err.Error())
		return
	}

	c.IndentedJSON(200, "")
}
