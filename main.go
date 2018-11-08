package main

import (
	"github.com/DiegoSantosWS/gowebscraping/config"
	"github.com/DiegoSantosWS/gowebscraping/helpers"
	"github.com/DiegoSantosWS/gowebscraping/model"
)

var dao = model.Connect{}
var conf = config.ConfigDB{}

func init() {
	conf.Read()

	dao.Server = conf.Server
	dao.Database = conf.Database
	dao.Connection()
}

func main() {
	helpers.WsiteBrasilBlog()
	helpers.UolNews()
	helpers.ExameNews()
}
