package app

import (
	"bookstore_users-api/logger"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	mapUrls()
	logger.Log.Info("about to start the application...")
	router.Run(":8080")
}
