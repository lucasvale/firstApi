package services

import (
	structs "data"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"persistence"
	"strings"
)

type EventServiceHandler struct {
	dbHandler persistence.DatabaseHandler
}

func NewHandlerEvent(databasehandler persistence.DatabaseHandler) *EventServiceHandler {
	return &EventServiceHandler{
		dbHandler: databasehandler,
	}
}

func (eh *EventServiceHandler) FindEventHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	criteria, ok := vars["criterioBusca"]
	if !ok {
		response.WriteHeader(404)
		fmt.Fprint(response, `{"error": "No search criteria found, you can either search by id via /id/4
							to search by name via /name/coldplayconcert}"`)
		return
	}

	searchkey, ok := vars["busca"]
	if !ok {
		response.WriteHeader(404)
		fmt.Fprint(response, `{"error": "No search keys found, you can either search by id via /id/4
							to search by name via /name/coldplayconcert}"`)
		return
	}

	var event structs.Event
	var err error
	switch strings.ToLower(criteria) {
	case "name":
		event, err = eh.dbHandler.FindEventByName(searchkey)
	case "id":
		id, err := hex.DecodeString(searchkey)
		if err == nil {
			event, err = eh.dbHandler.FindEvent(id)
		}

	}
	if err != nil {
		response.WriteHeader(404)
		fmt.Fprintf(response, `{"error": "%s"}`, err)
		return
	}

	response.Header().Set("Content-Type", "application/json;charset=utf8")
	json.NewEncoder(response).Encode(&event)
}
func (eh *EventServiceHandler) AllEventHandler(response http.ResponseWriter, request *http.Request) {
	events, err := eh.dbHandler.FindAllAvailableEvents()
	if err != nil {
		response.WriteHeader(500)
		fmt.Fprintf(response, "{error: Error occured while trying to find all available events %s}", err)
		return
	}
	response.Header().Set("Content-Type", "application/json;charset=utf8")
	err = json.NewEncoder(response).Encode(&events)
	if err != nil {
		response.WriteHeader(500)
		fmt.Fprintf(response, "{error: Error occured while trying encode events to JSON %s}", err)
	}
}

func (eh *EventServiceHandler) NewEventHandler(response http.ResponseWriter, request *http.Request) {
	event := structs.Event{}
	err := json.NewDecoder(request.Body).Decode(&event)

	if err != nil {
		response.WriteHeader(500)
		fmt.Fprint(response, "{error: error occured while decoding event data %s}", err)
		return
	}
	id, err := eh.dbHandler.AddEvent(event)
	if err != nil {
		response.WriteHeader(500)
		fmt.Fprint(response, "{error: error occured while persisting event %d %s}", id, err)
	}
}
