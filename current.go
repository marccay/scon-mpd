package main

import (
	"fmt"
)

func currentSong() {
	lines := getMap("currentsong")
	artist := lines["Artist"]
	if artist == "" {
		artist = "nil"
	}
	album := lines["Album"]
	if album == "" {
		album = "nil"
	}
	title := lines["Title"]
	if title == "" {
		title = "nil"
	}
	fmt.Println(artist + " : " + album + " - " + title)
}
