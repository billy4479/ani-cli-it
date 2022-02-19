package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

const (
	SEARCH_URL      = "https://www.animesaturn.it/animelist?search=%s"
	SEARCH_SELECTOR = "ul.list-group > li:nth-child(1) > div:nth-child(1) > div:nth-child(2) > h3:nth-child(1) > a:nth-child(1)"
)

func Search(name string) (*Askable, error) {
	res, err := http.Get(fmt.Sprintf(SEARCH_URL, url.QueryEscape(name)))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	results := []*Askable{}
	doc.Find(SEARCH_SELECTOR).Each(func(i int, s *goquery.Selection) {
		results = append(results, &Askable{Name: s.Text(), URL: s.AttrOr("href", "")})
	})
	if len(results) == 0 {
		color.New(color.FgRed).Printf("No results for %s\n", name)
		os.Exit(1)
	}

	i, err := MakeMenu(true, results)
	if err != nil {
		return nil, err
	}

	return results[i], nil
}
