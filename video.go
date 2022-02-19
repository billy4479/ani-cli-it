package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	PLAYER_URL_SELECTOR            = "div.container:nth-child(10) > div:nth-child(1) > div:nth-child(1) > a:nth-child(3)"
	VIDEO_URL_SELECTOR             = "source"
	VIDEO_URL_SELECTOR_ALTERNATIVE = ".embed-responsive-item > script:nth-child(4)"
)

func GetVideoURL(ep *Askable) (string, error) {
	res, err := http.Get(ep.URL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	return ExtractVideoUrlFromPlayer(doc.Find(PLAYER_URL_SELECTOR).AttrOr("href", ""))
}

func ExtractVideoUrlFromPlayer(playerUrl string) (string, error) {
	res, err := http.Get(playerUrl)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	if r := doc.Find(VIDEO_URL_SELECTOR); r.Length() != 0 {
		return r.AttrOr("src", ""), nil
	}

	js := strings.Split(doc.Find(VIDEO_URL_SELECTOR_ALTERNATIVE).Text(), "\n")
	for _, l := range js {
		if strings.Contains(l, "playlist.m3u8") {
			return strings.Split(l, "\"")[1], nil
		}
	}

	return "", fmt.Errorf("Unable to find video for %s, please report this error", playerUrl)
}
