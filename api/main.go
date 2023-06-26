package main

import (
	"log"
	"os"
	"products-api/api/controller"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func main() {
	sentryInit()

	app := gin.Default()
	app.Use(sentrygin.New(sentrygin.Options{}))
	productsRouting(app)
	categoriesRouting(app)
	picturesRouting(app)

	app.Run()
}

func productsRouting(router *gin.Engine) {
	router.GET("/products/:productID", controller.GetProductByID)
	router.GET("/products/category/:categoryID", controller.GetProductsByCategoryID)
	router.POST("/products", controller.AddProduct)
	router.PUT("/products", controller.UpdateProduct)
	router.DELETE("/products/:productID", controller.DeleteProduct)
}

func categoriesRouting(router *gin.Engine) {
	router.GET("/categories", controller.GetCategories)
	router.GET("/categories/:categoryID", controller.GetCategoryByID)
	router.POST("/categories", controller.AddCategory)
	router.PUT("/categories", controller.UpdateCategory)
	router.DELETE("/categories/:categoryID", controller.DeleteCategory)
}

func picturesRouting(router *gin.Engine) {
	router.GET("/pictures/:pictureID", controller.GetPictureByID)
	router.POST("/pictures", controller.UploadPicture)
	router.DELETE("/pictures/:pictureID", controller.DeletePicture)
}

func sentryInit() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	defer sentry.Flush(2 * time.Second)
}
