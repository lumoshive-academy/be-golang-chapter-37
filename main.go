package main

import (
	"golang-chapter-37/config"
	"golang-chapter-37/controller"
	"golang-chapter-37/database"
	"golang-chapter-37/middleware"
	"golang-chapter-37/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	config.LoadConfig()
}

func main() {
	r := gin.Default()

	// Middleware
	r.Use(middleware.Logger())
	r.Use(middleware.BasicAuth())

	controller := controller.NewController(database.GetDB())
	router.APIRouter(r, controller)

	r.Run(viper.GetString("SERVER_PORT"))
}
