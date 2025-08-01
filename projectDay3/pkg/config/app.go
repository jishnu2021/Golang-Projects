package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)

var (
	DB *gorm.DB
)

func Connect() {
	dsn := "root:Jishnu@2025@tcp(127.0.0.1:3306)/GolangTodo?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
