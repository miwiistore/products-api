package controller

import (
	"products-api/api/domain"
	"products-api/api/service"
	"strconv"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/miwiistore/utils/apierror"
)

func GetProductByID(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		apiError := apierror.NewBadRequestApiError("products", "Invalid parameter.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	product, err := service.GetProductByID(int64(productID))
	if err != nil {
		sentry.CaptureException(err)
		apiError := apierror.NewInternalServerApiError("products", "Error getting product.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	if product.ProductID == 0 {
		apiError := apierror.NewNotFoundApiError("products", "Product not found.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	c.IndentedJSON(200, product)
}

func GetProductsByCategoryID(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		apiError := apierror.NewBadRequestApiError("products", "Invalid parameter.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	products, err := service.GetProductsByCategoryID(int16(categoryID))
	if err != nil {
		sentry.CaptureException(err)
		apiError := apierror.NewInternalServerApiError("products", "Error getting products.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	if products == nil {
		apiError := apierror.NewNotFoundApiError("products", "Products not found.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	c.IndentedJSON(200, products)
}

func AddProduct(c *gin.Context) {
	var newProduct domain.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		apiError := apierror.NewBadRequestApiError("products", "Input body is invalid.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	newProduct, err := service.AddProduct(newProduct)
	if err != nil {
		sentry.CaptureException(err)
		apiError := apierror.NewInternalServerApiError("products", "Error creating product.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	c.IndentedJSON(200, newProduct)
}

func UpdateProduct(c *gin.Context) {
	var currentProduct domain.Product
	if err := c.ShouldBindJSON(&currentProduct); err != nil {
		apiError := apierror.NewBadRequestApiError("products", "Input body is invalid.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	if currentProduct.ProductID == 0 {
		apiError := apierror.NewBadRequestApiError("products", "Missing product_id. Field is required.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	currentProduct, err := service.UpdateProduct(currentProduct)
	if err != nil {
		sentry.CaptureException(err)
		apiError := apierror.NewInternalServerApiError("products", "Error updating product.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	c.IndentedJSON(200, currentProduct)
}

func DeleteProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		apiError := apierror.NewBadRequestApiError("products", "Invalid parameter.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	err = service.DeleteProduct(int64(productID))
	if err != nil {
		sentry.CaptureException(err)
		apiError := apierror.NewInternalServerApiError("products", "Error deleting product.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	c.IndentedJSON(200, "")
}
