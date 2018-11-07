package model

import (
	"log"

	"github.com/DiegoSantosWS/gowebscraping/stru"
	"gopkg.in/mgo.v2"
)

// Connect receive data to connection with database
type Connect struct {
	Server   string
	Database string
}

var db *mgo.Database

// Connection perform connection with database
func (m *Connect) Connection() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
		return
	}
	db = session.DB(m.Database)
}

// RegisterCollection register information in collection of the database connected
func (m *Connect) RegisterCollection(collection string, value stru.DataColecteds) error {
	err := db.C(collection).Insert(&value)
	return err
}
