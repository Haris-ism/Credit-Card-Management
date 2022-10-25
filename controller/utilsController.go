package controller

import (
	"gin/model"
	"log"

	"github.com/gin-gonic/gin"
)

func QueryFindCreditCards(t *Repo, c *gin.Context, condition string, value interface{}) (model.CreditCards, bool) {
	var creditCards model.CreditCards

	if err := t.DB.Find(&creditCards, condition, value).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(500, gin.H{
			"message": "failed to find users",
		})
		return creditCards, false
	}
	return creditCards, true
}
func QueryFindUsers(t *Repo, c *gin.Context, condition string, value interface{}) (model.Users, bool) {
	var user model.Users

	if err := t.DB.Find(&user, condition, value).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(500, gin.H{
			"message": "failed to find users",
		})
		return user, false
	}

	return user, true
}
func QueryFind(t *Repo, c *gin.Context, condition string, value string) (model.Accounts, bool) {
	var accounts model.Accounts

	if err := t.DB.Find(&accounts, condition, value).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(500, gin.H{
			"message": "failed to find email",
		})
		return accounts, false
	}

	return accounts, true
}
