package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase(){
	var err error
	godotenv.Load()
	dsn := os.Getenv("DB_NAME")
  	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		fmt.Println("Failed to Connect Database")
	}

	fmt.Println("Database Connected!")
}