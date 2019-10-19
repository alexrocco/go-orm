package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/alexrocco/go-orm/internal/models"
	"github.com/alexrocco/go-orm/internal/services"
	"github.com/gorilla/mux"
)

// PeopleHandler defines PeopleHander behaviours
type PeopleHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

// peopleHandlerImpl implements PeopleHandler interface
type peopleHandlerImpl struct {
	peopleService services.PeopleService
}

// NewPeopleHandler creates a new PeopleHandler
func NewPeopleHandler(service services.PeopleService) PeopleHandler {
	return peopleHandlerImpl{peopleService: service}
}

// Get a People resource
func (ph peopleHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	people := ph.peopleService.FindByID(id)

	if people.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	respondJSON(w, http.StatusOK, people)
}

// Update a People resource
func (ph peopleHandlerImpl) Update(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	people := models.People{}
	if err := decoder.Decode(&people); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if people.ID == 0 {
		respondError(w, http.StatusInternalServerError, "id not provided")
		return
	}

	ph.peopleService.SaveOrUpdate(people)

	w.WriteHeader(http.StatusNoContent)
}

// Create a People resource
func (ph peopleHandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	people := models.People{}
	if err := decoder.Decode(&people); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if people.ID > 0 {
		respondError(w, http.StatusInternalServerError, "id should be not provided")
		return
	}

	ph.peopleService.SaveOrUpdate(people)

	w.WriteHeader(http.StatusNoContent)
}

// Delete a People resource
func (ph peopleHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ph.peopleService.DeleteByID(id)

	w.WriteHeader(http.StatusNoContent)
}

// GetAll gets all the People resource
func (ph peopleHandlerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	people := ph.peopleService.FindAll()
	respondJSON(w, http.StatusOK, people)
}
