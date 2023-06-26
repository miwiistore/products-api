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
		apiError := apierror.NewInternalServerApiError("categories", "Error getting categories.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	if categories == nil {
		apiError := apierror.NewNotFoundApiError("categories", "Categories not found.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	c.IndentedJSON(200, categories)
}

func GetCategoryByID(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		apiError := apierror.NewBadRequestApiError("categories", "Invalid parameter.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	category, err := service.GetCategoryByID(int16(categoryID))
	if err != nil {
		sentry.CaptureException(err)
		apiError := apierror.NewInternalServerApiError("categories", "Error getting category.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	if category.CategoryID == 0 {
		apiError := apierror.NewNotFoundApiError("categories", "Category not found.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	c.IndentedJSON(200, category)
}

func AddCategory(c *gin.Context) {
	var newCategory domain.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		apiError := apierror.NewBadRequestApiError("categories", "Input body is invalid.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	newCategory, err := service.AddCategory(newCategory)
	if err != nil {
		sentry.CaptureException(err)
		apiError := apierror.NewInternalServerApiError("categories", "Error creating category.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	c.IndentedJSON(200, newCategory)
}

func UpdateCategory(c *gin.Context) {
	var currentCategory domain.Category
	if err := c.ShouldBindJSON(&currentCategory); err != nil {
		apiError := apierror.NewBadRequestApiError("categories", "Input body is invalid.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	if currentCategory.CategoryID == 0 {
		apiError := apierror.NewBadRequestApiError("categories", "Missing category_id. Field is required.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	currentCategory, err := service.UpdateCategory(currentCategory)
	if err != nil {
		sentry.CaptureException(err)
		apiError := apierror.NewInternalServerApiError("categories", "Error updating category.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	c.IndentedJSON(200, currentCategory)
}

func DeleteCategory(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		apiError := apierror.NewBadRequestApiError("categories", "Invalid parameter.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	err = service.DeleteCategory(int16(categoryID))
	if err != nil {
		sentry.CaptureException(err)
		apiError := apierror.NewInternalServerApiError("categories", "Error deleting category.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	c.IndentedJSON(200, "")
}
