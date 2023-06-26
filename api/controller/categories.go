package controller

import (
	"products-api/api/domain"
	"products-api/api/service"
	"strconv"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/miwiistore/utils/apierror"
)

func GetCategories(c *gin.Context) {
	categories, err := service.GetCategories()
	if err != nil {
		sentry.CaptureException(err)
		apiError := apierror.NewInternalServerApiError("products", "Error getting categories.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	if categories == nil {
		apiError := apierror.NewNotFoundApiError("products", "Categories not found.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	c.IndentedJSON(200, categories)
}

func GetCategoryByID(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		apiError := apierror.NewBadRequestApiError("products", "Invalid parameter.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	category, err := service.GetCategoryByID(int16(categoryID))
	if err != nil {
		sentry.CaptureException(err)
		apiError := apierror.NewInternalServerApiError("products", "Error getting category.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	if category.CategoryID == 0 {
		apiError := apierror.NewNotFoundApiError("products", "Category not found.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	c.IndentedJSON(200, category)
}

func AddCategory(c *gin.Context) {
	var newCategory domain.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		apiError := apierror.NewBadRequestApiError("products", "Input body is invalid.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	newCategory, err := service.AddCategory(newCategory)
	if err != nil {
		sentry.CaptureException(err)
		apiError := apierror.NewInternalServerApiError("products", "Error creating category.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	c.IndentedJSON(200, newCategory)
}

func UpdateCategory(c *gin.Context) {
	var currentCategory domain.Category
	if err := c.ShouldBindJSON(&currentCategory); err != nil {
		apiError := apierror.NewBadRequestApiError("products", "Input body is invalid.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	if currentCategory.CategoryID == 0 {
		apiError := apierror.NewBadRequestApiError("products", "Missing category_id. Field is required.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	currentCategory, err := service.UpdateCategory(currentCategory)
	if err != nil {
		sentry.CaptureException(err)
		apiError := apierror.NewInternalServerApiError("products", "Error updating category.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	c.IndentedJSON(200, currentCategory)
}

func DeleteCategory(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		apiError := apierror.NewBadRequestApiError("products", "Invalid parameter.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	err = service.DeleteCategory(int16(categoryID))
	if err != nil {
		sentry.CaptureException(err)
		apiError := apierror.NewInternalServerApiError("products", "Error deleting category.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	c.IndentedJSON(200, "")
}
