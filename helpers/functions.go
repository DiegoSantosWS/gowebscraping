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
		saveDatas(link, img, txt, "")
	})
}

// UolNews perform a scraping on the website uol news
func UolNews() {
	resp, err := http.Get("https://economia.uol.com.br/ultimas/")
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

	doc.Find(".thumbnail-standard").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Find("a").First().Attr("href")
		img, _ := s.Find("img").First().Attr("src")
		title := strings.TrimSpace(s.Find("h3").Text())
		timer := s.Find("time").Text()

		saveDatas(url, img, title, timer)
	})
}

// ExameNews perform a scraping on the website exame news
func ExameNews() {
	resp, err := http.Get("https://exame.abril.com.br/noticias-sobre/mercado-financeiro/")
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

	doc.Find(".list-item").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Find("a").First().Attr("href")
		img, _ := s.Find("img").First().Attr("src")
		title := strings.TrimSpace(s.Find(".list-item-title").Text())
		timer := s.Find(".list-date-description").Text()

		saveDatas(url, img, title, timer)
	})
}
