package service

import (
	"products-api/api/domain"
	"products-api/api/repository"
)

func GetCategories() ([]domain.Category, error) {
	return repository.GetCategories()
}

func GetCategoryByID(categoryID int16) (domain.Category, error) {
	return repository.GetCategoryByID(categoryID)
}

func AddCategory(newCategory domain.Category) (domain.Category, error) {
	return repository.AddCategory(newCategory)
}

func UpdateCategory(currentCategory domain.Category) (domain.Category, error) {
	return repository.UpdateCategory(currentCategory)
}

func DeleteCategory(categoryID int16) error {
	return repository.DeleteCategory(categoryID)
}
