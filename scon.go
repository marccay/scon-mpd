package main

import (
	"fmt"
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
		case "currentsong", "current", "song":
			currentSong()
		case "play":
			send("play")
		case "stop":
			send("stop")
		case "pause":
			send("pause")
		case "playpause":
			state := singleAtt("state")
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
			statusTree(opt)
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
		case "-h", "help", "--h":
			help()
		default:
			fmt.Println("unknown command, use [ scon-mpd -h ] for syntax help")
			os.Exit(1)
		}
	}

}

func help() {
	fmt.Println("usage:  scon-mpd [command] [option]")
	fmt.Println()
	fmt.Println("commands:")
	fmt.Println("  play\t\t play current song")
	fmt.Println("  pause\t\t toggles pause")
	fmt.Println("  playpause\t if stopped, play; if paused or playing toggles pause")
	fmt.Println("\t\t\t use with single button media keys")
	fmt.Println("  stop\t\t stop mpd")
	fmt.Println("  next\t\t play next song")
	fmt.Println("  previous\t play previous song")
	fmt.Println("\t option: smart | s")
	fmt.Println()
	fmt.Println("cli commands:")
	fmt.Println("  currentsong | current\t\t print info on current song")
	fmt.Println()
	fmt.Println("  repeat\t\t\t toggle repeat")
	fmt.Println("   : on | off")
	fmt.Println("  random\t\t\t toggle random")
	fmt.Println("   : on | off")
	fmt.Println("  single\t\t\t toggle single mode (repeat song)")
	fmt.Println("   : on | off")
	fmt.Println("  consume\t\t\t toggle consume (remove from playlist after play)")
	fmt.Println("   : on | off")
	fmt.Println("  crossfade | fade\t\t toggle crossfade (On or default 10 seconds)")
	fmt.Println("   : off | [seconds]")
	fmt.Println()
	fmt.Println("  status\t\t\t print status of toggles/controls")
	fmt.Println("   : state\t\t\t play state: play/pause/stopped")
	fmt.Println("   : repeat\t\t\t repeat toggle: on/off")
	fmt.Println("   : random | rand\t\t random toggle: on/off")
	fmt.Println("   : single\t\t\t single toggle: on/off")
	fmt.Println("   : consume\t\t\t consume toggle: on/off")
	fmt.Println("   : crossfade | fade\t\t crossfade toggle: off/[seconds]")
	fmt.Println("   : time\t\t\t prints elapsed/duration")
	fmt.Println("   : elapsed\t\t\t elapsed time of song in seconds")
	fmt.Println("   : duration\t\t\t total duration of song in seconds")

}
