package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

func Auth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		fmt.Println("failed to get cookies")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// fmt.Println("token", tokenString)
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT")), nil
	})
	// fmt.Println("token parsed", token)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("claims", claims["email"], claims["exp"])
		// fmt.Println("tes time", float64(time.Now().Add(time.Hour*24).Unix()))
		// fmt.Println("exp", claims["exp"])
		// var user model.Account

		// if err := handler.Service.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		// 	fmt.Println("failed to get data")
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"message": "email invalid",
		// 		"info":    err,
		// 	})
		// 	return
		// }
		// fmt.Println("user found", user)
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			fmt.Println("token expired euy")
			c.JSON(http.StatusForbidden, gin.H{
				"message": "token expired",
			})
		}
	} else {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Next()
}
