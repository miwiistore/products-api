package service

import (
	"products-api/api/domain"
	"products-api/api/repository"
	"products-api/api/storage"
)

func GetPictureByID(pictureID int64) (domain.Picture, error) {
	return repository.GetPictureByID(pictureID)
}

func UploadPictures(pictureBytes []byte, newPicture domain.Picture) (domain.Picture, error) {
	urlNewPicture, err := storage.UploadPicture(pictureBytes)
	if err != nil {
		return domain.Picture{}, err
	}
	newPicture.Url = urlNewPicture

	newPicture, err = repository.AddPicture(newPicture)
	if err != nil {
		return domain.Picture{}, err
	}

	return newPicture, nil
}

func DeletePicture(pictureID int64) error {
	picture, err := repository.GetPictureByID(pictureID)
	if err != nil {
		return err
	}

	if picture.PictureID == 0 {
		return nil
	}

	if err := storage.DeletePicture(picture.Url); err != nil {
		return err
	}

	if err := repository.DeletePicture(pictureID); err != nil {
		return err
	}

	return nil
}
