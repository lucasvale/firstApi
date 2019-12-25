package persistence

import "data"

type DatabaseHandler interface {
	AddEvent(structs.Event) ([]byte, error)
	FindEvent([]byte) (structs.Event, error)
	FindEventByName(string) (structs.Event, error)
	FindAllAvailableEvents() ([]structs.Event, error)
}
