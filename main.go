package main

import (
	"fmt"

	"gin/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	destination := "host=localhost user=postgres password=asdwasdw1 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(destination), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("db connected")
	}
	repo := handler.NewRepository(db)
	service := handler.NewService(repo)
	userHandler := handler.NewUserHandler(service)
	result, _ := service.FindAll()
	fmt.Println("result", result)
	users, _ := repo.FindAll()
	fmt.Println("ieu users", users)
	router := gin.Default()
	router.GET("/", userHandler.RootHandler)
	router.GET("/get", userHandler.Get)
	router.POST("/post", handler.Post)
	router.GET("/:id", handler.Params)
	router.GET("/query", handler.Queries)
	fmt.Println("server starts")
	router.Run(":6969")
}
