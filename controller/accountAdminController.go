package controller

import (
	"gin/model"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (t *Repo) SignUp(c *gin.Context) {
	var body model.Sign
	//method to get body of request
	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		return
	}
	accounts, bools := QueryFind(t, c, "email = ?", body.Email)
	if bools == false {
		return
	}
	if accounts.ID != 0 {
		log.Println("data not Found")
		c.JSON(400, gin.H{
			"message": "email is already used",
		})
		return
	}
	log.Println("query success euy")
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		log.Println("failed to hash password")
	}
	account := model.Accounts{}
	account.Email = body.Email
	account.Password = string(hash)
	// method to post to DB
	if err1 := t.DB.Create(&account).Error; err1 != nil {
		log.Println(err1)
		return
	}
	c.JSON(201, gin.H{
		"message": "write success",
		"data":    account,
	})
}

func (t *Repo) SignIn(c *gin.Context) {
	body := model.Sign{}
	//method to get body of request
	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		return
	}
	accounts, bools := QueryFind(t, c, "email = ?", body.Email)
	if bools == false {
		return
	}
	if accounts.ID == 0 {
		log.Println("data not Found")
		c.JSON(400, gin.H{
			"message": "invalid email",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(accounts.Password), []byte(body.Password)); err != nil {
		c.JSON(400, gin.H{
			"message": "invalid password",
		})
		return
	}
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": body.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT")))
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "/", "/", false, true)
	log.Println(tokenString, err)
	c.JSON(http.StatusOK, gin.H{
		"message": "Sign In Success",
		"data":    accounts,
		"token":   tokenString,
	})
}
