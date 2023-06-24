package service

import (
	"products-api/api/domain"
	"products-api/api/repository"
)

func GetProductByID(productID int64) (domain.Product, error) {
	return repository.GetProductByID(productID)
}

func GetProductsByCategoryID(categoryID int16) ([]domain.Product, error) {
	return repository.GetProductsByCategoryID(categoryID)
}

func AddProduct(newProduct domain.Product) (domain.Product, error) {
	return repository.AddProduct(newProduct)
}

func UpdateProduct(currentProduct domain.Product) (domain.Product, error) {
	return repository.UpdateProduct(currentProduct)
}

func DeleteProduct(productID int64) error {
	return repository.DeleteProduct(productID)
}
