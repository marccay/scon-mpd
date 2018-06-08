package main

import (
	"fmt"
	"os"
)

func statusTree(opt string) {
	switch opt {
	case "":
		printStatus()
	case "all":
		printStatus()
	case "random", "rand":
		fmt.Println(singleAtt("random"))
	case "repeat":
		fmt.Println(singleAtt(opt))
	case "single":
		fmt.Println(singleAtt(opt))
	case "consume":
		fmt.Println(singleAtt(opt))
	case "crossfade", "fade", "xfade":
		s := singleAtt("xfade")
		if s != "" {
			fmt.Println(singleAtt("xfade"))
		} else {
			fmt.Println("off")
		}
	case "state":
		fmt.Println(singleAtt(opt))
	case "time":
		fmt.Println(singleAtt(opt))
	case "elapsed":
		fmt.Println(singleAtt(opt))
	case "duration":
		fmt.Println(singleAtt(opt))
	default:
		fmt.Printf("status option [%v] not available\n", opt)
		os.Exit(1)
	}
}
