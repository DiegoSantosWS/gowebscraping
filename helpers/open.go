package helpers

import (
	"log"

	"github.com/DiegoSantosWS/gowebscraping/model"

	"github.com/DiegoSantosWS/gowebscraping/stru"
	"gopkg.in/mgo.v2/bson"
)

var m = model.Connect{}

// saveDatas save of data in collection
func saveDatas(url, img, desc, data string) bool {

	var d stru.DataColecteds
	d.ID = bson.NewObjectId()
	d.URL = url
	d.Image = img
	d.Description = desc
	d.Data = data

	if err := m.RegisterCollection("colects", d); err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
