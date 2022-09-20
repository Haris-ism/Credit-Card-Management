package handler

import (
	"fmt"
	"gin/model"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (t *Repo) Post(c *gin.Context) {
	body := model.BodyParser{}
	//method to get body of request
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("body:", body)

	user := model.User{}
	user.Name = body.Name
	user.Grade = body.Grade
	// method to post to DB
	if err1 := t.DB.Create(&user).Error; err1 != nil {
		fmt.Println(err1)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "write success",
		"data":    user,
	})

}
func (t *Repo) SignUp(c *gin.Context) {
	body := model.BodyParser{}
	//method to get body of request
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("body:", body)
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		fmt.Println("failed to hash password")
	}
	account := model.Account{}
	account.Email = body.Email
	account.Password = string(hash)
	// method to post to DB
	if err1 := t.DB.Create(&account).Error; err1 != nil {
		fmt.Println(err1)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "write success",
		"data":    account,
	})

}

func (t *Repo) SignIn(c *gin.Context) {
	body := model.BodyParser{}
	//method to get body of request
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("body:", body)

	// account := model.Account{}
	var user model.Account

	if err := t.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		fmt.Println("failed to get data")
		c.JSON(http.StatusOK, gin.H{
			"message": "email invalid",
			"info":    err,
		})
		return
	}
	if user.ID == 0 {
		fmt.Println("data not Found")
		c.JSON(http.StatusOK, gin.H{
			"message": "no data",
		})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "password invalid",
		})
		return
	}
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": body.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT")))
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	fmt.Println(tokenString, err)
	fmt.Println("getOne Data", user)
	c.JSON(http.StatusOK, gin.H{
		"message": "getOne success",
		"data":    user,
		"token":   tokenString,
	})

}
func (t *Repo) Get(c *gin.Context) {
	var user []model.User

	if err := t.DB.Find(&user).Error; err != nil {
		fmt.Println("failed to get data")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "get success",
		"data":    user,
	})
}
func (t *Repo) GetOne(c *gin.Context) {
	id := c.Param("id")
	var user []model.User

	if err := t.DB.Find(&user, id).Error; err != nil {
		fmt.Println("failed to get data")
		c.JSON(http.StatusOK, gin.H{
			"message": "failed to get data",
			"info":    err,
		})
		return
	}
	if len(user) == 0 {
		fmt.Println("data not Found")
		c.JSON(http.StatusOK, gin.H{
			"message": "no data",
		})
		return
	}
	fmt.Println("getOne Data", user)
	c.JSON(http.StatusOK, gin.H{
		"message": "getOne success",
		"data":    user,
	})
}
func (t *Repo) Delete(c *gin.Context) {
	id := c.Param("id")
	var user []model.User

	if err := t.DB.Delete(&user, id).Error; err != nil {
		fmt.Println("failed to delete data")
		c.JSON(http.StatusOK, gin.H{
			"message": "failed to delete data",
			"info":    err,
		})
		return
	}
	fmt.Println("delete success")
	c.JSON(http.StatusOK, gin.H{
		"message": "delete success",
	})
}
func (t *Repo) Put(c *gin.Context) {
	id := c.Param("id")
	body := model.BodyParser{}

	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		return
	}
	var user model.User

	if err := t.DB.Find(&user, id).Error; err != nil {
		fmt.Println(err)
		return
	}
	if user.ID == 0 {
		fmt.Println("no data")
		c.JSON(http.StatusOK, gin.H{
			"message": "no such data",
		})
		return
	}
	user.Name = body.Name
	user.Grade = body.Grade

	if err := t.DB.Save(&user).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "editing failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "edit success",
		"data":    user,
	})
}
