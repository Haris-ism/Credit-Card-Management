package main

import (
	"fmt"
	"net/http"
	"gin/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	// "reflect"
)
type trans struct {
	DB *gorm.DB	//this struct is required to make method for the established db connection
}
// func (t *trans)FindAll() ([]handler.User,error){
// 	var asd []handler.User
// 	err:=t.DB.Find(&asd).Error
// 	return asd,err
// }
func (t *trans)Post(c *gin.Context){
	body:=handler.BodyParser{}
	err:=c.BindJSON(&body) //method to get body of request
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("body:", body)
	}
	user:=handler.User{}
	user.Name=body.Name
	user.Grade=body.Grade
	err1:=t.DB.Create(&user).Error // method to post to DB
	if err1!=nil{
		fmt.Println(err1)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"read success",
		"data":user,
	})
	
}
func (t *trans)Get(c *gin.Context){
	var asd []handler.User
	err:=t.DB.Find(&asd).Error
	if err!=nil{
		fmt.Println("failed to get data")
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"write success",
		"data":asd,
	})
}
func (t *trans)GetOne(c *gin.Context){
	id:=c.Param("id")
	// fmt.Println("id",id,reflect.TypeOf(id))
	var asd []handler.User
	err:=t.DB.Find(&asd,id).Error
	if err!=nil{
		fmt.Println("failed to get data")
		c.JSON(http.StatusOK,gin.H{
			"message":"failed to get data",
			"info":err,
		})
		return
	}
	if len(asd)==0{
		fmt.Println("data not Found")
		c.JSON(http.StatusOK,gin.H{
			"message":"no data",
		})
		return
	}
	fmt.Println("getOne Data",asd)
	c.JSON(http.StatusOK,gin.H{
		"message":"getOne success",
		"data":asd,
	})
}
func (t *trans)Delete(c *gin.Context){
	id:=c.Param("id")
	var asd []handler.User
	err:=t.DB.Delete(&asd,id).Error
	if err!=nil{
		fmt.Println("failed to delete data")
		c.JSON(http.StatusOK,gin.H{
			"message":"failed to delete data",
			"info":err,
		})
		return
	}
	fmt.Println("delete success")
		c.JSON(http.StatusOK,gin.H{
			"message":"delete success",
		})
}
func (t *trans)Put(c *gin.Context){
	id:=c.Param("id")
	body:=handler.BodyParser{}
	err:=c.BindJSON(&body)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Print("body",body)
	var asd handler.User
	err1:=t.DB.Find(&asd,id).Error
	if err1!=nil{
		fmt.Println(err)
		return
	}
	if asd.ID==0{
		fmt.Println("no data")
		c.JSON(http.StatusOK,gin.H{
			"message":"no such data",
		})
		return
	}
	fmt.Println("data:",asd)
	asd.Name=body.Name
	asd.Grade=body.Grade
	err2:=t.DB.Save(&asd).Error
	if err2!=nil{
		fmt.Println(err)
		c.JSON(http.StatusOK,gin.H{
			"message":"editing failed",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"edit success",
		"data":asd,
	})
}
func service(db *gorm.DB) *trans{
	return &trans{db} // function to pass the db connection
}
func main(){
	dest:="host=localhost user=postgres password=asdwasdw1 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db,err:=gorm.Open(postgres.Open(dest),&gorm.Config{})
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("db connected")
	
	services:=service(db)
	router:=gin.Default()
	router.GET("/",services.Get)
	router.GET("/:id",services.GetOne)
	router.DELETE("/:id",services.Delete)
	router.PUT("/:id",services.Put)
	router.POST("/",services.Post)
	fmt.Println("starts")
	router.Run(":6969")
}