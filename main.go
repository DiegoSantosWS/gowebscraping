package main

import (
	"fmt"

	"github.com/DiegoSantosWS/gowebscraping/config"
	"github.com/DiegoSantosWS/gowebscraping/helpers"
	"github.com/DiegoSantosWS/gowebscraping/model"
	"github.com/jasonlvhit/gocron"
)

var dao = model.Connect{}
var conf = config.ConfigDB{}

func init() {
	conf.Read()

	dao.Server = conf.Server
	dao.Database = conf.Database
	dao.Connection()
}

func roundFunctions() {
	helpers.WsiteBrasilBlog()
	helpers.UolNews()
	helpers.ExameNews()
	helpers.UolEconomy()
	helpers.InfoMoney()
}
func main() {
	gocron.Every(1).Second().Do(helpers.WsiteBrasilBlog)
	gocron.Every(1).Second().Do(helpers.UolNews)
	gocron.Every(1).Second().Do(helpers.ExameNews)
	gocron.Every(1).Second().Do(helpers.UolEconomy)
	gocron.Every(1).Second().Do(helpers.InfoMoney)
	_, time := gocron.NextRun()
	fmt.Println(time)

	//gocron.Remove(roundFunctions)
	//gocron.Clear()

	<-gocron.Start()
}
