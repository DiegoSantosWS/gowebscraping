package model

import (
	"log"

	"github.com/DiegoSantosWS/gowebscraping/stru"
	"github.com/globalsign/mgo/bson"
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
func (m *Connect) RegisterCollection(collection string, value stru.DataColecteds) (err error) {
	if value.Reference == 0 {
		err = db.C(collection).Insert(&value)
		if err != nil {
			return err
		}
	} else {
		result := m.checkRegister(collection, value.Reference)

		if result == 0 {
			err = db.C(collection).Insert(&value)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// CheckRegister verify register to not insert duplicty of the information in collection of the database connected
func (m *Connect) checkRegister(collection string, reference int64) int64 {
	value := stru.DataColecteds{}

	err := db.C(collection).Find(bson.M{"cod_referency": reference}).Select(bson.M{"cod_referency": reference}).One(&value)
	if err != nil {
		return 0
	}

	return int64(value.Reference)
}
