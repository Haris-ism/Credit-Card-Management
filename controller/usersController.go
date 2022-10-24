package controller

import (
	"gin/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *Repo) Get(c *gin.Context) {
	var user []model.Users
	var getAll []model.UsersJoinCreditCards

	query := "users.id, users.name,users.job, credit_cards.credit_card_number,credit_cards.ammount,credit_cards.limit,credit_cards.bank,credit_cards.updated_at"
	join := "left join credit_cards on credit_cards.users_id = users.id and credit_cards.deleted_at is null"

	if err := t.DB.Model(&user).Select(query).Joins(join).Scan(&getAll).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(500, gin.H{
			"message": "failed to get data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "get success",
		"data":    getAll,
	})
}

func (t *Repo) GetOne(c *gin.Context) {
	id := c.Param("id")
	user := QueryFindUsers(t, c, "id = ?", id)
	if user == false {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "getOne success",
		"data":    user,
	})
}

func (t *Repo) Post(c *gin.Context) {
	var body model.InputData
	//method to get body of request
	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "failed to get body data",
		})
		return
	}

	var user model.Users
	user.Name = body.Name
	user.Job = body.Job

	// method to post to DB
	if err1 := t.DB.Create(&user).Error; err1 != nil {
		log.Println(err1)
		c.JSON(500, gin.H{
			"message": "failed to create data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "write success",
		"data":    user,
	})
}

func (t *Repo) Put(c *gin.Context) {
	id := c.Param("id")
	var body model.InputData
	var user model.Users

	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "failed to get body data",
		})
		return
	}
	// user := QueryFindUsers(t, c, "id = ?", body.UsersID)

	if err := t.DB.Find(&user, id).Error; err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "failed to find users",
		})
		return
	}
	if user.ID == 0 {
		log.Println("no data")
		c.JSON(400, gin.H{
			"message": "user id not found",
		})
		return
	}
	user.Name = body.Name
	user.Job = body.Job
	if err := t.DB.Save(&user).Error; err != nil {
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

func (t *Repo) Delete(c *gin.Context) {
	id := c.Param("id")
	user := QueryFindUsers(t, c, "id = ?", id)
	if user == false {
		return
	}
	if err := t.DB.Delete(&user, id).Error; err != nil {
		log.Println("users delete failed :", err)
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
