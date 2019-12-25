package dblayer

import (
	"persistence"
	"persistence/mongoLayer"
)

type DBTYPE string

const (
	MONGODB  DBTYPE = "mongodb"
	DYNAMODB DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {
	switch options {
	case MONGODB:
		return mongoLayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}
