package app

import (
	"log"
	"net/http"

	handler "github.com/alexrocco/go-orm/internal/app/go-orm/handlers"
	"github.com/gorilla/mux"
)

type App struct {
	router *mux.Router
	peopleHandler handler.PeopleHandler
}

const peopleRequestPath  = "/people"

func NewApp(peopleHandler handler.PeopleHandler) (app App) {
	app = App{
		router: mux.NewRouter(),
		peopleHandler: peopleHandler,
	}
	app.setRouters()

	return
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.router))
}

func (a *App) setRouters() {
	// People requests
	a.get(peopleRequestPath, a.handleRequest(a.peopleHandler.GetAll))
	a.get(peopleRequestPath + "/{id:[0-9]+}", a.handleRequest(a.peopleHandler.Get))
	a.post(peopleRequestPath, a.handleRequest(a.peopleHandler.Create))
	a.put(peopleRequestPath, a.handleRequest(a.peopleHandler.Update))
	a.delete(peopleRequestPath + "/{id:[0-9]+}", a.handleRequest(a.peopleHandler.Delete))
}

type requestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler requestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}

// get wraps the router for GET method
func (a *App) get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, f).Methods("GET")
}

// post wraps the router for POST method
func (a *App) post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, f).Methods("POST")
}

// put wraps the router for PUT method
func (a *App) put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, f).Methods("PUT")
}

// delete wraps the router for DELETE method
func (a *App) delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, f).Methods("DELETE")
}
