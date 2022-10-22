package controller

import (
	"gin/model"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// var asd int

func (t *Repo) RegistrationCC(c *gin.Context) {
	body := model.InputCreditCard{}
	var user model.Users
	//method to get body of request
	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		return
	}

	if err := t.DB.Find(&user, body.UsersID).Error; err != nil {
		log.Println("failed to find users id", err)
		c.JSON(http.StatusOK, gin.H{
			"message": "failed to find users id",
		})
		return
	}
	if user.ID == 0 {
		log.Println("no data")
		c.JSON(http.StatusOK, gin.H{
			"message": "users id not found",
		})
		return
	}
	var creditCards model.CreditCards
	creditCards.UsersID = body.UsersID
	creditCards.Bank = body.Bank
	creditCards.Limit = body.Limit

	// method to post to DB
	if err := t.DB.Create(&creditCards).Error; err != nil {
		log.Println("error when registering cc to the database euy", err)
		c.JSON(http.StatusOK, gin.H{
			"message": "failed to regist credit card",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "write success",
		"data":    body,
	})
}
func (t *Repo) Post(c *gin.Context) {
	body := model.InputData{}
	//method to get body of request
	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		return
	}

	var user model.Users
	user.Name = body.Name
	user.Job = body.Job

	// method to post to DB
	if err1 := t.DB.Create(&user).Error; err1 != nil {
		log.Println(err1)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "write success",
		"data":    user,
	})

}
func (t *Repo) SignUp(c *gin.Context) {
	body := model.Sign{}
	//method to get body of request
	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		return
	}
	log.Println("body:", body)
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
	c.JSON(http.StatusOK, gin.H{
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
	log.Println("body:", body)

	var user model.Accounts

	if err := t.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(http.StatusOK, gin.H{
			"message": "email is invalid",
			"info":    err,
		})
		return
	}
	if user.ID == 0 {
		log.Println("data not Found")
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

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT")))
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "/", "/", false, true)
	log.Println(tokenString, err)
	log.Println("getOne Data", user)
	c.JSON(http.StatusOK, gin.H{
		"message": "getOne success",
		"data":    user,
		"token":   tokenString,
	})

}
func (t *Repo) Get(c *gin.Context) {
	var user []model.Users
	var getAll []model.UsersJoinCreditCards

	if err := t.DB.Model(&user).Select("users.id, users.name,users.job, credit_cards.credit_card_number,credit_cards.ammount,credit_cards.limit,credit_cards.bank,credit_cards.updated_at").Joins("left join credit_cards on credit_cards.users_id = users.id").Scan(&getAll).Error; err != nil {
		log.Println("failed to get data")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "get success",
		"data":    getAll,
	})
}
func (t *Repo) GetOne(c *gin.Context) {
	id := c.Param("id")
	var user []model.Users

	if err := t.DB.Find(&user, id).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(http.StatusOK, gin.H{
			"message": "failed to get data",
			"info":    err,
		})
		return
	}
	if len(user) == 0 {
		log.Println("data not Found")
		c.JSON(http.StatusOK, gin.H{
			"message": "no data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "getOne success",
		"data":    user,
	})
}
func (t *Repo) Delete(c *gin.Context) {
	id := c.Param("id")
	var user []model.Users

	if err := t.DB.Delete(&user, id).Error; err != nil {
		log.Println("failed to delete data")
		c.JSON(http.StatusOK, gin.H{
			"message": "failed to delete data",
			"info":    err,
		})
		return
	}
	log.Println("delete success")
	c.JSON(http.StatusOK, gin.H{
		"message": "delete success",
	})
}

func (t *Repo) UpdateCreditCards(c *gin.Context) {
	body := model.InputCreditCard{}
	var user model.Users
	var creditCard model.CreditCards

	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		return
	}
	if err := t.DB.Find(&user, body.UsersID).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "failed to search users id",
		})
		return
	}
	if user.ID == 0 {
		log.Println("no data")
		c.JSON(http.StatusOK, gin.H{
			"message": "user id not found",
		})
		return
	}
	if err := t.DB.Find(&creditCard, "users_id = ?", body.UsersID).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "failed to search credit card id",
		})
		return
	}
	if creditCard.ID == 0 {
		log.Println("no data")
		c.JSON(http.StatusOK, gin.H{
			"message": "credit card is not registered",
		})
		return
	}
	creditCard.UsersID = body.UsersID
	creditCard.Bank = body.Bank
	creditCard.Limit = body.Limit
	creditCard.Ammount = body.Ammount

	if err := t.DB.Save(&creditCard).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "update failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
		"data":    user,
	})
}
func (t *Repo) Put(c *gin.Context) {
	id := c.Param("id")
	body := model.InputData{}

	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		return
	}
	var user model.Users
	log.Println("ieu put user", reflect.TypeOf(id))
	if err := t.DB.Find(&user, id).Error; err != nil {
		log.Println(err)
		return
	}
	if user.ID == 0 {
		log.Println("no data")
		c.JSON(http.StatusOK, gin.H{
			"message": "user id not found",
		})
		return
	}
	user.Name = body.Name
	user.Job = body.Job
	if err := t.DB.Save(&user).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "update failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
		"data":    user,
	})
}
