package model

import (
	"time"

	"gorm.io/gorm"
)

type BodyParser struct {
	Name     string `json:"name"`
	Grade    int    `json:"grade"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Account struct {
	gorm.Model
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	gorm.Model
	Name       string    `json:"name" binding:"required"`
	Grade      int       `json:"grade" binding:"required"`
	Created_At time.Time `json:"created_at"`
	IeuEmail   string
	Account    Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:IeuEmail;references:Email"`
}
