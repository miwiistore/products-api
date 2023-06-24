package repository

import (
	"products-api/api/domain"
)

func GetCategories() ([]domain.Category, error) {
	resultCategories, err := clientDB.Query(
		"SELECT CATEGORY_ID, NAME FROM products.CATEGORIES")
	if err != nil {
		return nil, err
	}

	var categories []domain.Category
	for resultCategories.Next() {
		var category domain.Category

		err = resultCategories.Scan(&category.CategoryID, &category.Name)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func GetCategoryByID(categoryID int16) (domain.Category, error) {
	resultCategories, err := clientDB.Query(
		"SELECT CATEGORY_ID, NAME FROM products.CATEGORIES WHERE CATEGORY_ID = ?",
		categoryID,
	)
	if err != nil {
		return domain.Category{}, err
	}

	var category domain.Category
	for resultCategories.Next() {
		err = resultCategories.Scan(&category.CategoryID, &category.Name)
		if err != nil {
			return domain.Category{}, err
		}
	}

	return category, nil
}

func AddCategory(newCategory domain.Category) (domain.Category, error) {
	newRecord, err := clientDB.Exec(
		"INSERT INTO products.CATEGORIES (NAME) VALUES(?)",
		newCategory.Name,
	)
	if err != nil {
		return domain.Category{}, err
	}

	categoryID, err := newRecord.LastInsertId()
	if err != nil {
		return domain.Category{}, err
	}
	newCategory.CategoryID = int16(categoryID)

	return newCategory, nil
}

func UpdateCategory(currentCategory domain.Category) (domain.Category, error) {
	_, err := clientDB.Exec(
		"UPDATE products.CATEGORIES SET NAME = ? WHERE CATEGORY_ID = ?",
		currentCategory.Name,
		currentCategory.CategoryID,
	)
	if err != nil {
		return domain.Category{}, err
	}

	return currentCategory, nil
}

func DeleteCategory(categoryID int16) error {
	_, err := clientDB.Exec(
		"DELETE FROM products.CATEGORIES WHERE CATEGORY_ID = ?",
		categoryID,
	)
	if err != nil {
		return err
	}

	return nil
}
