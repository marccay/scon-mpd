package main

import (
	"fmt"
)

func current_song() {
	lines := get_map("currentsong")
	artist := lines["Artist"]
	album := lines["Album"]
	title := lines["Title"]
	fmt.Println(artist + " : " + album + " - " + title)
}
