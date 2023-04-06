package connection

import (
	"github.com/Bappy60/BookStore_in_Go/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "root:@/bookstore_db?charset=utf8&parseTime=True&loc=Local")
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
	db.AutoMigrate(&models.Book{},&models.Author{})
	return db
}


