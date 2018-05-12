package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		os.Exit(1)
	} else {
		cmd := os.Args[1]
		var opt string
		if len(os.Args) == 3 {
			opt = os.Args[2]
		}
		switch cmd {
			case "play":
				send("play")
			case "stop":
				send("stop")
			case "pause":
				send("pause")
			case "playpause":
				state := single_att("state")
				if state == "play" {
					send("pause")
				} else if state == "pause" {
					send("pause")
				} else {
					send("play")
				}
			case "next":
				send("next")
			case "previous":
				previous(opt)
			case "status":
				status_tree(opt)
			case "repeat":
				toggle("repeat", opt)
			case "random", "rand":
				toggle("random", opt)
			case "single":
				toggle("single", opt)
			case "consume":
				toggle("consume", opt)
			case "crossfade", "xfade", "fade":
				crossfade(opt)
			default:
				os.Exit(1)
		}
	}

}

