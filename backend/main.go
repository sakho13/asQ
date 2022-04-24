package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sakho13/backend/api"
)

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

	engine.Run(":8080")
}
