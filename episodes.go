package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

const (
	EPISODE_LIST_QUERY = "#range-anime-0"
)

func GetEpisode(sr *Askable) (*Askable, error) {
	res, err := http.Get(sr.URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	results := []*Askable{}
	doc.Find(EPISODE_LIST_QUERY).Children().Each(func(i int, s *goquery.Selection) {
		results = append(results, &Askable{Name: strings.TrimSpace(s.Text()), URL: s.Children().AttrOr("href", "")})
	})
	if len(results) == 0 {
		color.New(color.FgRed).Printf("No episodes for %s\n", sr.Name)
		os.Exit(1)
	}

	i, err := MakeMenu(true, results)
	if err != nil {
		return nil, err
	}

	return results[i], nil

}
