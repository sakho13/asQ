package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sakho13/backend/types"
)

func Hello(c *gin.Context) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		c.JSON(http.StatusOK, types.CommonResponseType[interface{}]{
			ResultFlg: 0,
			Message:   err.Error(),
			Response:  nil,
		})
		return
	}
	now := time.Now().In(jst)

	response := types.CommonResponseType[string]{
		ResultFlg: 1,
		Message:   "Hello!!",
		Response:  now.Format("2006/01/02 15:04:05"),
	}

	c.JSON(http.StatusOK, response)
}
