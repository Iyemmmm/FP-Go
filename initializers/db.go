package initializers

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/fp_store?charset=utf8mb4&parseTime=True&loc=Local"
  	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		fmt.Println("Failed to Connect Database")
	}

	fmt.Println("Database Connected!")
}