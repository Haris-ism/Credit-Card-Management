package main

import (
	"fmt"
	"gin/handler"
	"gin/initial"
	"gin/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	initial.LoadEnv()
	port := os.Getenv("PORT")
	db := initial.ConnectDB()
	services := handler.Service(db)
	router := gin.Default()
	router.Use(middleware.CorsMiddleware)
	router.GET("/", middleware.Auth, services.Get)
	router.GET("/:id", middleware.Auth, services.GetOne)
	router.DELETE("/:id", middleware.Auth, services.Delete)
	router.PUT("/:id", middleware.Auth, services.Put)
	router.POST("/", middleware.Auth, services.Post)
	router.POST("/signup", services.SignUp)
	router.POST("/signin", services.SignIn)
	router.POST("/goroutines", services.Goroutines)
	fmt.Println("starts")
	router.Run(port)
}
