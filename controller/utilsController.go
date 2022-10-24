package controller

import (
	"gin/model"
	"log"

	"github.com/gin-gonic/gin"
)

func QueryFindUsers(t *Repo, c *gin.Context, condition string, value interface{}) interface{} {
	var user model.Users

	if err := t.DB.Find(&user, condition, value).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(500, gin.H{
			"message": "failed to find users",
		})
		return false
	}
	if user.ID == 0 {
		log.Println("data not Found")
		c.JSON(500, gin.H{
			"message": "users not found",
		})
		return false
	}
	return user
}
func QueryFind(t *Repo, c *gin.Context, condition string, value string) interface{} {
	var user model.Accounts

	// var user interface{}
	// user=model.Accounts{}
	// var asd interface{}
	// asd="aaa"
	// log.Println("type asd",reflect.TypeOf(asd))
	// asd=1
	// log.Println("type asd",reflect.TypeOf(asd))
	// asd=model.Accounts{}
	// log.Println("type asd",reflect.TypeOf(asd))
	// log.Println("type user",user.(model.Accounts).Email)
	// // log.Println("type user1",reflect.TypeOf(user1.Email),user1.Email)

	// log.Println("ieu find na ",condition,value)
	// switch condition{
	// 	case "id = ?":
	// 	user=model.Users{}
	// }

	if err := t.DB.Find(&user, condition, value).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(500, gin.H{
			"message": "failed to find email",
		})
		return false
	}
	if user.ID != 0 {
		log.Println("data not Found")
		c.JSON(500, gin.H{
			"message": "email is already used",
		})
		return false
	}
	return true
}
