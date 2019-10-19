package repositories

import (
	"github.com/alexrocco/go-orm/internal/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Import sqlite3
)

// NewDB creates a new DB connection
func NewDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", ":memory:")
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
