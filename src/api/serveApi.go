package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"persistence"
	"services"
)

func ServeApi(endpoint string, dbHandler persistence.DatabaseHandler) error {
	handler := services.NewHandlerEvent(dbHandler)
	r := mux.NewRouter()
	eventsRouter := r.PathPrefix("/events").Subrouter()
	eventsRouter.Methods("GET").Path("/{criterioBusca}/{busca}").HandlerFunc(handler.FindEventHandler)
	eventsRouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)
	eventsRouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler)
	return http.ListenAndServe(endpoint, r)
}
