package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	type responseType struct {
		Message string `json:"message"`
		From    string `json:"from"`
		Now     string `json:"now"`
	}

	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(jst)

	response := responseType{
		Message: "Hello!?",
		From:    "WithMe",
		Now:     now.Format("2006/01/02 15:04:05"),
	}

	c.JSON(http.StatusOK, response)
}
