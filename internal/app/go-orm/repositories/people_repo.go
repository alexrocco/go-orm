package repositories

import (
	"github.com/alexrocco/go-orm/internal/app/go-orm/models"
	"github.com/jinzhu/gorm"
)

type PeopleRepo struct {
	db *gorm.DB
}

func NewPeopleRepo(db *gorm.DB) PeopleRepo{
	return PeopleRepo{db}
}

func (p PeopleRepo) FindAll() []models.People {
	var people []models.People
	p.db.Find(&people)
	return people
}

func (p PeopleRepo) FindById(id int) models.People {
	var people models.People
	p.db.First(&people, id)
	return people
}

func (p PeopleRepo) DeleteById(id int) {
	var people models.People
	p.db.First(&people, id)
	p.db.Delete(&people)
}

func (p PeopleRepo) SaveOrUpdate(people models.People) {
	p.db.Save(people)
}
