package handlers

import (
	"encoding/json"
	"github.com/alexrocco/go-orm/internal/app/go-orm/models"
	"github.com/alexrocco/go-orm/internal/app/go-orm/services"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// PeopleHandler holds services used in the struct
type PeopleHandler struct {
	peopleService services.PeopleService
}

// NewPeopleHandler creates a new PeopleHandler
func NewPeopleHandler(service services.PeopleService) PeopleHandler {
	return PeopleHandler{peopleService: service}
}

// Get a People resource
func (ph PeopleHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	people := ph.peopleService.FindById(id)

	if people.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	respondJSON(w, http.StatusOK, people)
}

// Update a People resource
func (ph PeopleHandler) Update(w http.ResponseWriter, r *http.Request) {
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
func (ph PeopleHandler) Create(w http.ResponseWriter, r *http.Request) {
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
func (ph PeopleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ph.peopleService.DeleteById(id)

	w.WriteHeader(http.StatusNoContent)
}

// GetAll gets all the People resource
func (ph PeopleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	people := ph.peopleService.FindAll()
	respondJSON(w, http.StatusOK, people)
}