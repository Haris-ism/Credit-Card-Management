package controller

import (
	"gin/middleware"

	"github.com/gin-gonic/gin"
)

// var asd int
func MainRouter(services *Repo, port string) {
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
	routes.DELETE("/creditcards/:id", services.DeleteCC)

	router.Run(port)
}
