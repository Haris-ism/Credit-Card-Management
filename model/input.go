package model

import (
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
	Name  string `json:"name" binding:"required"`
	Grade int    `json:"grade" binding:"required"`
	// IeuEmail string
	// Account  Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:IeuEmail;references:Email"`
}
type CreditCard struct {
	gorm.Model
	Number     string
	UserNumber string
}

// type Usera struct {
// 	gorm.Model
// 	MemberNumber string       `gorm:"unique"`
// 	CreditCards  []CreditCard `gorm:"foreignKey:UserNumber;references:MemberNumber"`
// }
