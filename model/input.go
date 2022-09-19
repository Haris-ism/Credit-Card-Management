package model
import (
	"time"
)
type BodyParser struct {
	Name     string `json:"name"`
	Grade    int    `json:"grade"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name" binding:"required"`
	Grade      int    `json:"grade" binding:"required"`
	Created_At time.Time `json:"created_at"`
}
type Account struct {
	ID       int    `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
