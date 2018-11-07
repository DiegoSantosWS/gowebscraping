package helpers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Infomoney() {
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
		link, ret := s.Find("a").First().Attr("href")
		if ret {
			fmt.Println("Link: ", strings.TrimSpace(link))
		}
		img, rets := s.Find("img").First().Attr("data-pagespeed-lazy-src")
		if rets {
			fmt.Println("Img: ", strings.TrimSpace(img))
		}

		fmt.Println("Descrição: ", strings.TrimSpace(s.Find("h3").Text()))

	})
}
