package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sakho13/backend/api"
	"github.com/sakho13/backend/types"
)

func init() {
	api.DBInit()
	api.InitFirebase()
}

func main() {
	engine := gin.Default()

	// jst, _ := time.LoadLocation("Asia/Tokyo")
	// logFile, _ := os.Create("./logs/" + time.Now().In(jst).Format("2006-01-02_15-04-05") + ".log")
	// gin.DefaultWriter = io.MultiWriter(os.Stdout, logFile)

	// cors対応
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:8891",
			"http://localhost:8890",
		},
		AllowMethods: []string{
			"POST",
			"GET",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	// v1 API
	v1 := engine.Group("/v1")
	{
		v1.GET("/hello", api.Hello)

		v1Api := v1.Group("/api")
		v1Api.Use(commonMiddleware())
		{
			v1Api.GET("/user", api.GetUser)
			v1Api.POST("/user", api.CreateUser)
			v1Api.PUT("/user", api.EditUser)
		}
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

func commonMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("PATH: %v\n", ctx.FullPath())
		start := time.Now()

		ctx.Next()

		latency := time.Since(start)
		log.Printf("Latency: %v", latency)
	}
}
