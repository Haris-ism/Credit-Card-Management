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

// var asd int
func QueryFind(t *Repo,c *gin.Context,condition string,value string) interface{}{
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

	if err := t.DB.Find(&user,condition, value).Error; err != nil {
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
func (t *Repo) SignUp(c *gin.Context) {
	// var user model.Accounts
	body := model.Sign{}
	//method to get body of request
	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		return
	}
	if QueryFind(t,c,"email = ?",body.Email)==false{
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
	var user model.Accounts
	//method to get body of request
	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		return
	}

	if QueryFind(t,c,"email = ?",body.Email)==false{
		return
	}
	log.Println("query success euy")

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(400, gin.H{
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
	c.JSON(http.StatusOK, gin.H{
		"message": "Sign In Success",
		"data":    user,
		"token":   tokenString,
	})
}

func (t *Repo) Get(c *gin.Context) {
	var user []model.Users
	var getAll []model.UsersJoinCreditCards

	query:="users.id, users.name,users.job, credit_cards.credit_card_number,credit_cards.ammount,credit_cards.limit,credit_cards.bank,credit_cards.updated_at"
	join:="left join credit_cards on credit_cards.users_id = users.id and credit_cards.deleted_at is null"
	
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
	var user model.Users

	if err := t.DB.Find(&user,"id = ?", id).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(500, gin.H{
			"message": "failed to get data",
		})
		return
	}
	if user.ID == 0 {
		log.Println("data not Found")
		c.JSON(400, gin.H{
			"message": "no data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "getOne success",
		"data":    user,
	})
}

func (t *Repo) Post(c *gin.Context) {
	body := model.InputData{}
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

func (t *Repo) RegistrationCC(c *gin.Context) {
	body := model.InputCreditCard{}
	var user model.Users
	//method to get body of request
	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "failed to get body data",
		})
		return
	}

	if err := t.DB.Find(&user, body.UsersID).Error; err != nil {
		log.Println("failed to find users id", err)
		c.JSON(500, gin.H{
			"message": "failed to find users id",
		})
		return
	}
	if user.ID == 0 {
		log.Println("register cc: users id not found")
		c.JSON(400, gin.H{
			"message": "users id not found",
		})
		return
	}
	var creditCards model.CreditCards
	if err := t.DB.Find(&creditCards,"users_id = ?",user.ID).Error; err != nil {
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

func (t *Repo) Put(c *gin.Context) {
	id := c.Param("id")
	body := model.InputData{}
	var user model.Users

	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "failed to get body data",
		})
		return
	}
	
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
		c.JSON(500, gin.H{
			"message": "failed to search users id",
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

func (t *Repo) Delete(c *gin.Context) {
	id := c.Param("id")
	var user model.Users
	if err := t.DB.Find(&user, id).Error; err != nil {
		log.Println("delete cc failed :",err)
		c.JSON(500, gin.H{
			"message": "failed to get data",
		})
		return
	}
	if user.ID == 0 {
		log.Println("users id not Found")
		c.JSON(400, gin.H{
			"message": "users id not found",
		})
		return
	}
	if err := t.DB.Delete(&user, id).Error; err != nil {
		log.Println("users delete failed :",err)
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

func (t *Repo) DeleteCC(c *gin.Context) {
	id := c.Param("id")
	var creditCards model.CreditCards
	
	if err := t.DB.Find(&creditCards,"users_id = ?", id).Error; err != nil {
		log.Println("delete cc failed :",err)
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
		log.Println("failed to delete data: ",err)
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


