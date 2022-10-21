package main

import (
	"gin/controller"
	"gin/initial"
	"gin/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	initial.LoadEnv()
	port := os.Getenv("PORT")
	db := initial.ConnectDB()
	services := controller.Service(db)
	router := gin.Default()
	router.Use(middleware.CorsMiddleware)
	router.POST("/signup", services.SignUp)
	router.POST("/signin", services.SignIn)

	routes := router.Use(middleware.Auth)

	routes.GET("/", services.Get)
	routes.GET("/:id", services.GetOne)
	routes.DELETE("/:id", services.Delete)
	routes.PUT("/:id", services.Put)
	routes.POST("/", services.Post)
	routes.POST("/creditcards", services.RegistrationCC)
	routes.PUT("/creditcards", services.UpdateCreditCards)

	router.Run(port)
	log.Println("server starts")
}
