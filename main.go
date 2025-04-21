package main

import (
	"github.com/bot-trap/backend/config"
	"github.com/bot-trap/backend/models"
	"github.com/bot-trap/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.Post{})

	r := gin.Default()
}
