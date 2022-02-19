package main

import (
	"log"
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

	episode, err := GetEpisode(search)
	if err != nil {
		return err
	}

	videoUrl, err := GetVideoURL(episode)
	if err != nil {
		return err
	}

	return PlayURL(videoUrl)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
