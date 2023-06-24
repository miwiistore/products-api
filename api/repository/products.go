package repository

import (
	"products-api/api/domain"
)

func GetProductByID(productID int64) (domain.Product, error) {
	results, err := clientDB.Query(
		"SELECT PRODUCT_ID, CATEGORY_ID, TITLE, DESCRIPTION, QUANTITY, PRICE, SKU FROM products.PRODUCTS WHERE PRODUCT_ID = ?",
		productID,
	)
	if err != nil {
		return domain.Product{}, err
	}

	var product domain.Product
	for results.Next() {
		err = results.Scan(
			&product.ProductID,
			&product.CategoryID,
			&product.Title,
			&product.Description,
			&product.Quantity,
			&product.Price,
			&product.SKU,
		)
		if err != nil {
			return domain.Product{}, err
		}
	}

	return product, nil
}

func GetProductsByCategoryID(categoryID int16) ([]domain.Product, error) {
	results, err := clientDB.Query(
		"SELECT PRODUCT_ID, CATEGORY_ID, TITLE, DESCRIPTION, QUANTITY, PRICE, SKU FROM products.PRODUCTS WHERE CATEGORY_ID = ?",
		categoryID,
	)
	if err != nil {
		return nil, err
	}

	var products []domain.Product
	for results.Next() {
		var product domain.Product

		err = results.Scan(
			&product.ProductID,
			&product.CategoryID,
			&product.Title,
			&product.Description,
			&product.Quantity,
			&product.Price,
			&product.SKU,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func AddProduct(newProduct domain.Product) (domain.Product, error) {
	newRecord, err := clientDB.Exec(
		"INSERT INTO products.PRODUCTS (CATEGORY_ID, TITLE, DESCRIPTION, QUANTITY, PRICE, SKU) VALUES(?, ?, ?, ?, ?, ?)",
		newProduct.CategoryID,
		newProduct.Title,
		newProduct.Description,
		newProduct.Quantity,
		newProduct.Price,
		newProduct.SKU,
	)
	if err != nil {
		return domain.Product{}, err
	}

	productID, err := newRecord.LastInsertId()
	if err != nil {
		return domain.Product{}, err
	}
	newProduct.ProductID = int64(productID)

	return newProduct, nil
}

func UpdateProduct(currentProduct domain.Product) (domain.Product, error) {
	_, err := clientDB.Exec(
		"UPDATE products.PRODUCTS SET CATEGORY_ID = ?, TITLE = ?, DESCRIPTION = ?, QUANTITY = ?, PRICE = ?, SKU = ? WHERE PRODUCT_ID = ?",
		currentProduct.CategoryID,
		currentProduct.Title,
		currentProduct.Description,
		currentProduct.Quantity,
		currentProduct.Price,
		currentProduct.SKU,
		currentProduct.ProductID,
	)
	if err != nil {
		return domain.Product{}, err
	}

	return currentProduct, nil
}

func DeleteProduct(productID int64) error {
	_, err := clientDB.Exec(
		"DELETE FROM products.PRODUCTS WHERE PRODUCT_ID = ?",
		productID,
	)
	if err != nil {
		return err
	}

	return nil
}
