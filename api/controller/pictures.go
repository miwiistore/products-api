package controller

import (
	"encoding/base64"
	"products-api/api/domain"
	"products-api/api/service"
	"strconv"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/miwiistore/utils/apierror"
)

const (
	DOMAIN_NAME = "pictures"
)

func UploadPicture(c *gin.Context) {
	var newPicture domain.Picture
	if err := c.ShouldBindJSON(&newPicture); err != nil {
		apiError := apierror.NewBadRequestApiError(DOMAIN_NAME, "Input body is invalid.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	pictureBytes, err := base64.StdEncoding.DecodeString(newPicture.Base64)
	if err != nil {
		sentry.CaptureException(err)
		apiError := apierror.NewInternalServerApiError(DOMAIN_NAME, "Error encoding picture to PNG.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	pictures, err := service.UploadPictures(pictureBytes, newPicture)
	if err != nil {
		sentry.CaptureException(err)
		apiError := apierror.NewInternalServerApiError(DOMAIN_NAME, "Error storing pictures.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	c.IndentedJSON(200, pictures)
}

func DeletePicture(c *gin.Context) {
	pictureID, err := strconv.Atoi(c.Param("pictureID"))
	if err != nil {
		apiError := apierror.NewBadRequestApiError("categories", "Invalid parameter. Picture ID is required.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	err = service.DeletePicture(int64(pictureID))
	if err != nil {
		sentry.CaptureException(err)
		apiError := apierror.NewInternalServerApiError(DOMAIN_NAME, "Error deleting picture.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	c.IndentedJSON(200, "")
}

func GetPictureByID(c *gin.Context) {
	pictureID, err := strconv.Atoi(c.Param("pictureID"))
	if err != nil {
		apiError := apierror.NewBadRequestApiError(DOMAIN_NAME, "Invalid parameter.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	picture, err := service.GetPictureByID(int64(pictureID))
	if err != nil {
		sentry.CaptureException(err)
		apiError := apierror.NewInternalServerApiError(DOMAIN_NAME, "Error getting picture.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	if picture.PictureID == 0 {
		apiError := apierror.NewNotFoundApiError(DOMAIN_NAME, "Picture not found.")
		c.IndentedJSON(apiError.Status(), apiError)
		return
	}

	c.IndentedJSON(200, picture)
}
