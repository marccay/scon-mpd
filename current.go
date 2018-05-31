package main

import (
	"fmt"
)

func currentSong() {
	lines := getMap("currentsong")
	artist := lines["Artist"]
	album := lines["Album"]
	title := lines["Title"]
	fmt.Println(artist + " : " + album + " - " + title)
}
