package middleware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

const recaptchaURL = "https://www.recaptcha.net/api/siteverify"

type recaptchaResponse struct {
	Success bool `json:"success"`
}

func RecaptchaCheckMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var secretKey = secretKey
		response := c.GetHeader("G-Token")
		if response == "" {
			c.JSON(401, "Seems you're a robot :)")
			c.Abort()
			return
		}
		remoteIP := c.GetHeader("X-Forwarded-IP")
		resp, err := http.PostForm(recaptchaURL, url.Values{
			"secret":   {secretKey},
			"response": {response},
			"remoteip": {remoteIP},
		})
		if err != nil {
			c.JSON(500, "Error in connecting google recaptcha server")
			c.Abort()
			return
		}
		body, _ := ioutil.ReadAll(resp.Body)
		var res recaptchaResponse
		json.Unmarshal(body, &res)
		if !res.Success {
			c.JSON(401, "Seems you're a robot :(")
			c.Abort()
			return
		}
	}
}
