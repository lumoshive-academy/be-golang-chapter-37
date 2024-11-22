package main

import (
	"golang-chapter-37/implem-gin-gonic/config"
	"golang-chapter-37/implem-gin-gonic/controller"
	"golang-chapter-37/implem-gin-gonic/database"
	"golang-chapter-37/implem-gin-gonic/middleware"
	"golang-chapter-37/implem-gin-gonic/router"

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
