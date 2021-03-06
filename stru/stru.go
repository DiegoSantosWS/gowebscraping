package stru

import "gopkg.in/mgo.v2/bson"

// DataColecteds receive data for insert
type DataColecteds struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	URL         string        `bson:"link" json:"link"`
	Image       string        `bson:"img,omitempty" json:"img,omitempty"`
	Description string        `bson:"description" json:"description"`
	Data        string        `bson:"data,omitempty" json:"data,omitempty"`
	Reference   int64         `bson:"cod_referency,omitempty" json:"referency,omitempty"`
}
