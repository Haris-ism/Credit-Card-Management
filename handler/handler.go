package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService Service
}

func NewUserHandler(userService Service) *UserHandler {
	return &UserHandler{userService}
}
func (handler *UserHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "bambang",
		"info": "this is get method",
	})
}

func (handler *UserHandler) Get(c *gin.Context) {
	// destination := "host=localhost user=postgres password=asdwasdw1 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(destination), &gorm.Config{})
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("db connected")
	// }
	// user := []User{}

	// err = db.Find(&user).Error
	result, err := handler.userService.FindAll()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("read db success")
		fmt.Println(result)
		c.JSON(http.StatusOK, gin.H{
			"message": "read success",
			"data":    result,
		})
	}
}

func Post(c *gin.Context) {
	body := BodyParser{}
	err := c.BindJSON(&body)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("body:", body)
	}
	destination := "host=localhost user=postgres password=asdwasdw1 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(destination), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("db connected")
	}
	user := User{}
	user.Name = body.Name
	user.Grade = body.Grade

	err = db.Create(&user).Error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("write db success")
		c.JSON(http.StatusOK, gin.H{
			"message": "write success",
			"data":    user,
		})
	}

}
func Params(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
func Queries(c *gin.Context) {
	query := c.Query("query")
	c.JSON(http.StatusOK, gin.H{
		"query": query,
	})
}
