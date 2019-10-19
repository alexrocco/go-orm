package repositories

import (
	"github.com/alexrocco/go-orm/internal/models"
	"github.com/jinzhu/gorm"
)

// PeopleRepo defines PeopleRepo behaviours
type PeopleRepo interface {
	FindAll() []models.People
	FindByID(id int) models.People
	DeleteByID(id int)
	SaveOrUpdate(people models.People)
}

// PeopleRepoImpl holds dependencies
type PeopleRepoImpl struct {
	db *gorm.DB
}

// NewPeopleRepo creates a new PeopleRepo with database connection
func NewPeopleRepo(db *gorm.DB) PeopleRepo {
	return PeopleRepoImpl{db}
}

// FindAll People resource
func (p PeopleRepoImpl) FindAll() []models.People {
	var people []models.People
	p.db.Find(&people)
	return people
}

// FindByID People resource
func (p PeopleRepoImpl) FindByID(id int) models.People {
	var people models.People
	p.db.First(&people, id)
	return people
}

// DeleteByID People resource
func (p PeopleRepoImpl) DeleteByID(id int) {
	var people models.People
	p.db.First(&people, id)
	p.db.Delete(&people)
}

// SaveOrUpdate People resource
func (p PeopleRepoImpl) SaveOrUpdate(people models.People) {
	p.db.Save(people)
}
