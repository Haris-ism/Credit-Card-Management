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
	router.GET("/:id", services.GetOne)
	router.DELETE("/:id", services.Delete)
	router.PUT("/:id", services.Put)
	router.POST("/", services.Post)
	router.POST("/signup", services.SignUp)
	router.POST("/signin", services.SignIn)
	fmt.Println("starts")
	router.Run(port)
}
