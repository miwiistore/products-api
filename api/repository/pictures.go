package repository

import "products-api/api/domain"

func GetPictureByID(pictureID int64) (domain.Picture, error) {
	results, err := clientDB.Query(
		"SELECT PICTURE_ID, PRODUCT_ID, URL, `ORDER` FROM products.PICTURES WHERE PICTURE_ID = ?",
		pictureID,
	)
	if err != nil {
		return domain.Picture{}, err
	}

	var picture domain.Picture
	for results.Next() {
		err = results.Scan(
			&picture.PictureID,
			&picture.ProductID,
			&picture.Url,
			&picture.Order,
		)
		if err != nil {
			return domain.Picture{}, err
		}
	}

	return picture, nil
}

func AddPicture(newPicture domain.Picture) (domain.Picture, error) {
	newRecord, err := clientDB.Exec(
		"INSERT INTO products.PICTURES (PRODUCT_ID, URL, `ORDER`) VALUES(?, ?, ?)",
		newPicture.ProductID,
		newPicture.Url,
		newPicture.Order,
	)
	if err != nil {
		return domain.Picture{}, err
	}

	pictureID, err := newRecord.LastInsertId()
	if err != nil {
		return domain.Picture{}, err
	}
	newPicture.ProductID = int64(pictureID)

	return newPicture, nil
}

func DeletePicture(pictureID int64) error {
	_, err := clientDB.Exec(
		"DELETE FROM products.PICTURES WHERE PICTURE_ID = ?",
		pictureID,
	)
	if err != nil {
		return err
	}

	return nil
}
