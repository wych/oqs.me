package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"oqs.me/logger"
	"oqs.me/services"
	"oqs.me/utils"
)

func RegisterURL(c *gin.Context) {
	realURL := c.PostForm("url")
	urlRecord := &services.URLRecord{RealURL: realURL}
	userRecord := &services.UserRecord{}
	userRecord.IP = c.GetHeader("X-Real-IP")
	userRecord.ForwardedFor = c.GetHeader("X-Forwarded-For")
	userRecord.UA = c.GetHeader("User-Agent")
	id, err := services.IssueID()
	if err != nil {
		logger.ErrorLogger.Printf("Error when redis issue record id: %s", err.Error())
		c.JSON(500, "Server internal Error")
		return
	}
	err = urlRecord.CreateAndCache(id)
	if err != nil {
		logger.ErrorLogger.Printf("Error when database create record: %s", err.Error())
		c.JSON(500, "Server internal Error")
		return
	}
	userRecord.OQSRecord.ID = id
	err = userRecord.Save()
	if err != nil {
		logger.ErrorLogger.Printf("Error when database create user record: %s", err.Error())
	}
	c.JSON(200, string(urlRecord.ShortURL))

}

func ConvertURL(c *gin.Context) {
	shortURL := c.Param("shortURL")
	urlRecord := &services.URLRecord{ShortURL: utils.Base62Int(shortURL)}
	if !urlRecord.Valid() {
		c.String(400, "Bad request")
		return
	}
	if urlRecord.GetCache() {
		c.Redirect(301, urlRecord.RealURL)
	}
	err := urlRecord.FetchAndCache()
	if err != nil {
		log.Print(err)
	}
	c.Redirect(301, urlRecord.RealURL)
}
