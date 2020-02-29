package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"oqs.me/config"
	"oqs.me/controllers"
	"oqs.me/logger"
	"oqs.me/models"
	"oqs.me/services"
)

var recaptchaSecertKey string

func init() {
	config.Conf.InitConfFromToml("config.toml")
	logger.LoggerInit()
	models.DBInit()
	services.RedisInit()
	id, err := models.GetMaxOQSID()
	if err != nil {
		log.Fatalf("Error in init next record id: %s", err)
	}
	err = services.SetOQSID(id)
	if err != nil {
		log.Fatalf("Error in set next record id: %s", err)
	}
}

func main() {
	c := gin.Default()
	// recaptchaMiddleWare := middleware.RecaptchaCheckMiddleware(config.Conf.General.RecaptchaSecretKey)
	c.POST("/", controllers.RegisterURL)
	c.GET("/:shortURL", controllers.ConvertURL)

	c.Run(fmt.Sprintf("%s:%d", config.Conf.General.Listen, config.Conf.General.Port))
}
