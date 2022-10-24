package controller

import (
	"gin/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *Repo) RegistrationCC(c *gin.Context) {
	var body model.InputCreditCard
	//method to get body of request
	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "failed to get body data",
		})
		return
	}
	user := QueryFindUsers(t, c, "id = ?", body.UsersID)
	if user == false {
		return
	}
	var creditCards model.CreditCards
	if err := t.DB.Find(&creditCards, "users_id = ?", user.(model.Users).ID).Error; err != nil {
		log.Println("failed to find users id", err)
		c.JSON(500, gin.H{
			"message": "failed to find users id",
		})
		return
	}
	if creditCards.UsersID != 0 {
		log.Println("register cc: credit card already registered before")
		c.JSON(400, gin.H{
			"message": "credit card already registered before",
		})
		return
	}
	creditCards.UsersID = body.UsersID
	creditCards.Bank = body.Bank
	creditCards.Limit = body.Limit
	// method to post to DB
	if err := t.DB.Create(&creditCards).Error; err != nil {
		log.Println("error when registering cc to the database euy", err)
		c.JSON(500, gin.H{
			"message": "failed to regist credit card",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "write success",
		"data":    body,
	})
}

func (t *Repo) UpdateCreditCards(c *gin.Context) {
	var body model.InputCreditCard
	var creditCard model.CreditCards

	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		return
	}

	user := QueryFindUsers(t, c, "id = ?", body.UsersID)
	if user == false {
		return
	}
	if err := t.DB.Find(&creditCard, "users_id = ?", body.UsersID).Error; err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "failed to search credit card id",
		})
		return
	}
	if creditCard.ID == 0 {
		log.Println("no data")
		c.JSON(400, gin.H{
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
		c.JSON(500, gin.H{
			"message": "update failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "update success",
		"data":    user,
	})
}

func (t *Repo) DeleteCC(c *gin.Context) {
	id := c.Param("id")
	var creditCards model.CreditCards

	if err := t.DB.Find(&creditCards, "users_id = ?", id).Error; err != nil {
		log.Println("delete cc failed :", err)
		c.JSON(500, gin.H{
			"message": "failed to get data",
		})
		return
	}

	if creditCards.UsersID == 0 {
		log.Println("credit card not Found")
		c.JSON(400, gin.H{
			"message": "credit card not Found",
		})
		return
	}
	if err := t.DB.Delete(&creditCards, "users_id = ?", id).Error; err != nil {
		log.Println("failed to delete data: ", err)
		c.JSON(500, gin.H{
			"message": "failed to delete data",
		})
		return
	}

	log.Println("delete success")
	c.JSON(http.StatusOK, gin.H{
		"message": "delete success",
	})
}
