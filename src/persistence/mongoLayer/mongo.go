package mongoLayer

import (
	structs "data"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"persistence"
)

const (
	DB     = "myevents"
	USERS  = "users"
	EVENTS = "events"
)

type mongoDBLayer struct {
	session *mgo.Session
}

func (mgoLayer *mongoDBLayer) getFreshSession() *mgo.Session {
	return mgoLayer.session.Copy()
}

func (mgoLayer mongoDBLayer) AddEvent(e structs.Event) ([]byte, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}

	if !e.Location.ID.Valid() {
		e.Location.ID = bson.NewObjectId()
	}

	return []byte(e.ID), s.DB(DB).C(EVENTS).Insert(e)
}

func (mgoLayer mongoDBLayer) FindEvent(id []byte) (structs.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := structs.Event{}
	err := s.DB(DB).C(EVENTS).FindId(bson.ObjectId(id)).One(&e)
	return e, err
}

func (mgoLayer mongoDBLayer) FindAllAvailableEvents() ([]structs.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	var events []structs.Event
	err := s.DB(DB).C(EVENTS).Find(nil).All(&events)
	return events, err
}

func (mgoLayer mongoDBLayer) FindEventByName(name string) (structs.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := structs.Event{}
	err := s.DB(DB).C(EVENTS).Find(bson.M{"name": name}).One(&e)
	return e, err
}

func NewMongoDBLayer(connection string) (persistence.DatabaseHandler, error) {
	s, err := mgo.Dial(connection)
	return mongoDBLayer{
		session: s,
	}, err
}
