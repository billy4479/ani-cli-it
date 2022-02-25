package main

import (
	"log"
	"strings"
)

func run() error {
	name, err := StringOption("Search", nil)
	if err != nil {
		return err
	}

	search, err := Search(name)
	if err != nil {
		return err
	}

	playAnotherOne := true
	for playAnotherOne {
		episode, err := GetEpisode(search)
		if err != nil {
			return err
		}

		videoUrl, err := GetVideoURL(episode)
		if err != nil {
			return err
		}

		err = PlayURL(videoUrl)
		if err != nil {
			return err
		}

		r, err := StringOption("Play another episode? [Y/n]", func(s string) bool { return true })
		if err != nil {
			return err
		}

		if strings.ToLower(r) == "n" {
			break
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
