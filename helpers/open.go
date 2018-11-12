package helpers

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/DiegoSantosWS/gowebscraping/model"

	"github.com/DiegoSantosWS/gowebscraping/stru"
	"gopkg.in/mgo.v2/bson"
)

var m = model.Connect{}

// saveDatas save of data in collection
func saveDatas(collection, url, img, desc, data string, reference int) bool {

	var d stru.DataColecteds
	d.ID = bson.NewObjectId()
	d.URL = url
	d.Image = img
	d.Description = desc
	d.Data = data
	d.Reference = int64(reference)

	if err := m.RegisterCollection(collection, d); err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// ExtractID extract code to reference of the news colletcting
func ExtractID(str string) (int, error) {
	re := regexp.MustCompile(`\/(\d+)\/`)
	dts := re.FindStringSubmatch(str)
	cod := fmt.Sprintf("%s", dts[1]) //Convert array to string

	codI, err := strconv.Atoi(cod)
	if err != nil {
		return 0, err
	}

	return codI, nil
}

// ExtractCode extract code to reference of the news colletcting
func ExtractCode(str, params string, position int) int {

	u, err := url.Parse(str)
	if err != nil {
		log.Fatal(err)
	}
	uri := u.RequestURI()
	res := strings.Split(uri, params)
	codI, err := strconv.Atoi(res[position])
	if err != nil {
		log.Fatal("", err)
		return 0
	}
	return codI
}
