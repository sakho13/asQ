package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sakho13/backend/api"
	"github.com/sakho13/backend/types"
)

func init() {
	api.DBInit()
}

func main() {
	engine := gin.Default()

	// jst, _ := time.LoadLocation("Asia/Tokyo")
	// logFile, _ := os.Create("./logs/" + time.Now().In(jst).Format("2006-01-02_15-04-05") + ".log")
	// gin.DefaultWriter = io.MultiWriter(os.Stdout, logFile)

	// v1 API
	v1 := engine.Group("/v1")
	{
		v1.GET("/hello", api.Hello)

		api := v1.Group("/api")
		api.GET("/")
	}

	// when received route not found
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, types.CommonResponseType[interface{}]{
			ResultFlg: 0,
			Message:   "The route couldn't found.",
			Response:  nil,
		})
	})

	engine.Run(":8080")
}
