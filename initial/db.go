package initial

import (
	"fmt"
	"gin/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dbUrl := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("db connected")
	db.AutoMigrate(&model.User{}, &model.Account{})
	return db
}
