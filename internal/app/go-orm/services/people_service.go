package services

import (
	"github.com/alexrocco/go-orm/internal/app/go-orm/models"
	"github.com/alexrocco/go-orm/internal/app/go-orm/repositories"
)

type PeopleService struct {
	peopleRepo repositories.PeopleRepo
}

func NewPeopleService(repo repositories.PeopleRepo) PeopleService {
	return PeopleService{peopleRepo:repo}
}

func (s PeopleService) FindAll() []models.People {
	return s.peopleRepo.FindAll()
}

func (s PeopleService) FindById(id int) models.People {
	return s.peopleRepo.FindById(id)
}

func (s PeopleService) DeleteById(id int) {
	s.peopleRepo.DeleteById(id)
}

func (s PeopleService) SaveOrUpdate(people models.People) {
	s.peopleRepo.SaveOrUpdate(people)
}