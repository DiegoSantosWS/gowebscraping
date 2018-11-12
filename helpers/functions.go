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
		cod := ExtractCode(link, "/", 5)
		img, _ := s.Find("img").First().Attr("src")

		if strings.TrimSpace(img) == "/pagespeed_static/1.JiBnMqyl6S.gif" {
			img, _ = s.Find("img").First().Attr("data-pagespeed-lazy-src")
		}
		txt := strings.TrimSpace(s.Find("h3").Text())
		saveDatas("wsitebrasil", link, img, txt, "", cod)
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

		saveDatas("uol_last", url, img, title, timer, 0)
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

		saveDatas("abril", url, img, title, timer, 0)
	})
}

// UolEconomy
func UolEconomy() {
	resp, err := http.Get("https://economia.uol.com.br/")
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
		timer := s.Find(".thumb-time").Text()
		saveDatas("uol_economy", url, img, title, timer, 0)
	})
}

// InfoMoney
func InfoMoney() {
	res, err := http.Get("https://www.infomoney.com.br/ultimas-noticias")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}
	body, err := goquery.NewDocumentFromReader(res.Body)
	body.Find(".column").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Find("a").First().Attr("href")
		link := "https://www.infomoney.com.br" + url
		img, _ := s.Find("img").First().Attr("src")
		title, _ := s.Find("img").First().Attr("title")
		timer, _ := s.Find("i").First().Attr("class")

		if url != "" && img != "/assets/images/logo/infomoney-brown.png?v=2018b" {
			cod, _ := ExtractID(string(link))
			saveDatas("infomoney", link, img, title, timer, cod)
		}

	})
}
