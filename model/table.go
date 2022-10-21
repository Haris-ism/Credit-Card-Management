package model

import (
	"time"

	"gorm.io/gorm"
)

type Accounts struct {
	gorm.Model
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

type Users struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	Job  string `json:"job,omitempty" binding:"required"`
}
type CreditCards struct {
	gorm.Model
	UpdatedAt        time.Time `json:"updated_at"`
	CreditCardNumber int       `gorm:"autoIncrement"json:"credit_card_number" `
	Bank             string    `json:"bank"`
	Ammount          int       `json:"ammount"`
	Limit            int       `json:"limit"`
	UsersID          int
	Users            Users `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
type UsersJoinCreditCards struct {
	ID               int    `json:"id"`
	Name             string `json:"name,omitempty"`
	Job              string `json:"job,omitempty"`
	CreditCardNumber int    `json:"credit_card_number,omitempty"`
	Bank             string `json:"bank,omitempty"`
	Ammount          int    `json:"ammount,omitempty"`
	Limit            int    `json:"limit,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
}
