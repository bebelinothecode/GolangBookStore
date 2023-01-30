package config

import (
	// "os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	// "github.com/joho/godotenv"
)

var db *gorm.DB;



func Connect()  {
	// dbpassword := os.Getenv("DB_PASS");
	d, err := gorm.Open("mysql","isaac:password@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local");
	if err != nil {
		panic(err)
	}
	db = d;
}

func GetDB() *gorm.DB {
	return db;
}