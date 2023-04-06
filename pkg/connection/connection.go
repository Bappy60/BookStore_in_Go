package connection

import (
	"fmt"

	"github.com/Bappy60/BookStore_in_Go/pkg/config"
	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Connect() {
	config := config.GConfig
	connectionString := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
	config.DBUser, config.DBPass, config.DBIP, config.DbName)
	d, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	DB = d
}
func GetDB() *gorm.DB {
	return DB
}

func Initialize() *gorm.DB {
	Connect()
	db := GetDB()
	db.AutoMigrate(&models.Book{}, &models.Author{})
	return db
}
