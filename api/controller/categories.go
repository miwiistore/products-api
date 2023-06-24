package controller

import (
	"products-api/api/domain"
	"products-api/api/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	categories, err := service.GetCategories()
	if err != nil {
		c.IndentedJSON(500, err.Error())
		return
	}

	c.IndentedJSON(200, categories)
}

func GetCategoryByID(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		c.IndentedJSON(400, "Invalid parameter.")
		return
	}

	categories, err := service.GetCategoryByID(int16(categoryID))
	if err != nil {
		c.IndentedJSON(500, err.Error())
		return
	}

	c.IndentedJSON(200, categories)
}

func AddCategory(c *gin.Context) {
	var newCategory domain.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.IndentedJSON(400, "Input body is invalid.")
		return
	}

	newCategory, err := service.AddCategory(newCategory)
	if err != nil {
		c.IndentedJSON(500, err.Error())
		return
	}

	c.IndentedJSON(200, newCategory)
}

func UpdateCategory(c *gin.Context) {
	var currentCategory domain.Category
	if err := c.ShouldBindJSON(&currentCategory); err != nil {
		c.IndentedJSON(400, "Input body is invalid.")
		return
	}

	if currentCategory.CategoryID == 0 {
		c.IndentedJSON(400, "Missing category_id. Field is required.")
		return
	}

	currentCategory, err := service.UpdateCategory(currentCategory)
	if err != nil {
		c.IndentedJSON(500, err.Error())
		return
	}

	c.IndentedJSON(200, currentCategory)
}

func DeleteCategory(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		c.IndentedJSON(400, "Invalid parameter.")
		return
	}

	err = service.DeleteCategory(int16(categoryID))
	if err != nil {
		c.IndentedJSON(500, err.Error())
		return
	}

	c.IndentedJSON(200, "")
}
