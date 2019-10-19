package services

import (
	"github.com/alexrocco/go-orm/internal/models"
	"github.com/alexrocco/go-orm/internal/repositories"
)

// PeopleService defines PeopleService behaviours
type PeopleService interface {
	FindAll() []models.People
	FindByID(id int) models.People
	DeleteByID(id int)
	SaveOrUpdate(people models.People)
}

// peopleServiceImpl implements PeopleService
type peopleServiceImpl struct {
	peopleRepo repositories.PeopleRepo
}

// NewPeopleService creates a new PeopleService with its dependencies
func NewPeopleService(repo repositories.PeopleRepo) PeopleService {
	return peopleServiceImpl{peopleRepo: repo}
}

// FindAll People resource
func (s peopleServiceImpl) FindAll() []models.People {
	return s.peopleRepo.FindAll()
}

// FindById People resource
func (s peopleServiceImpl) FindByID(id int) models.People {
	return s.peopleRepo.FindByID(id)
}

// DeleteById People resource
func (s peopleServiceImpl) DeleteByID(id int) {
	s.peopleRepo.DeleteByID(id)
}

// SaveOrUpdate People resource
func (s peopleServiceImpl) SaveOrUpdate(people models.People) {
	s.peopleRepo.SaveOrUpdate(people)
}
