package stru

import "gopkg.in/mgo.v2/bson"

// DataColecteds receive data for insert
type DataColecteds struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	URL         string        `bson:"link" json:"link"`
	Image       string        `bson:"img" json:"img"`
	Description string        `bson:"description" json:"description"`
}
