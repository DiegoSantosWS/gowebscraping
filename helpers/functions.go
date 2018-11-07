package helpers

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// WsiteBrasilBlog perform a scraping on the website wsitebrasil
func WsiteBrasilBlog() {
	resp, err := http.Get("https://www.wsitebrasil.com.br/blog")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".box-news").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Find("a").First().Attr("href")

		img, _ := s.Find("img").First().Attr("src")

		if strings.TrimSpace(img) == "/pagespeed_static/1.JiBnMqyl6S.gif" {
			img, _ = s.Find("img").First().Attr("data-pagespeed-lazy-src")
		}
		txt := strings.TrimSpace(s.Find("h3").Text())
		saveDatas(link, img, txt)
	})
}

// UolNews perform a scraping on the website uol news
func UolNews() {

}

// GloboNews perform a scraping on the website g1
func GloboNews() {

}
