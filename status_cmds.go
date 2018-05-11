package main

import (
	"fmt"
	"os"
)


func status_tree (opt string) {
	switch opt {
		case "":
			print_status()
		case "all":
			print_status()
		case "random", "rand":
			fmt.Println(single_att("random"))
		case "repeat":
			fmt.Println(single_att(opt))
		case "single":
			fmt.Println(single_att(opt))
		case "consume":
			fmt.Println(single_att(opt))
		case "crossfade", "fade", "xfade":
			fmt.Println(single_att("xfade"))
		case "state":
			fmt.Println(single_att(opt))
		case "time":
			fmt.Println(single_att(opt))
		case "elapsed":
			fmt.Println(single_att(opt))
		case "duration":
			fmt.Println(single_att(opt))
		default:
			fmt.Println("option not available")
			os.Exit(1)
	}
}
