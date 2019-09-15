package repositories

import (
	"github.com/alexrocco/go-orm/internal/app/go-orm/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func NewDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:test123@tcp(localhost:3306)/goorm?charset=utf8")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.People{})

	db.Create(&models.People{
		FirstName: "Foo",
		LastName:  "Bar",
		Gender:    models.Male,
	})

	return db
}
