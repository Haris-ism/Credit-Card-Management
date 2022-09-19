package handler

import (
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB //this struct is required to make method for the established db connection
}

func Service(db *gorm.DB) *repo {
	return &repo{db} // function to pass the db connection
}
