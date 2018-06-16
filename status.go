package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

func singleAtt(att string) string {
	attributes := getMap("status")
	if att == "state" || att == "duration" || att == "elapsed" || att == "xfade" {
		return attributes[att]
	} else if att == "time" {
		return attributes["elapsed"] + "/" + attributes["duration"]
	} else {
		isOn, _ := isAttOn(attributes[att])
		return isOn
	}
}

func isAttOn(stat string) (string, int) {
	if stat == "0" {
		return "off", 0
	}
	return "on", 1
}

func songTime() (float64, float64, uint64) {
	att := getMap("status")
	duration, err := strconv.ParseFloat(att["duration"], 64)
	if err != nil {
		os.Exit(1)
	}
	elapsed, err := strconv.ParseFloat(att["elapsed"], 64)
	if err != nil {
		os.Exit(1)
	}
	songID, err := strconv.ParseUint(att["song"], 10, 64)
	if err != nil {
		os.Exit(1)
	}
	return elapsed, duration, songID

}

func printStatus() {
	att := getMap("status")
	fmt.Println("state:", att["state"])
	isOn, _ := isAttOn(att["repeat"])
	fmt.Println("repeat:", isOn)
	isOn, _ = isAttOn(att["random"])
	fmt.Println("random:", isOn)
	isOn, _ = isAttOn(att["single"])
	fmt.Println("single:", isOn)
	isOn, _ = isAttOn(att["consume"])
	fmt.Println("consume:", isOn)
	if att["xfade"] != "" {
		fmt.Println("crossfade:", att["xfade"])
	} else {
		fmt.Println("crossfade:", "off")
	}

	if att["state"] == "play" || att["state"] == "pause" {
		fmt.Println("elapsed", att["elapsed"])
		fmt.Println("duration", att["duration"])
	}
}

func getMap(val string) map[string]string {
	var address string
	addr, err := readConfig()
	if err != nil {
		address = ":6600"
	} else {
		address = addr
	}

	conn, err := net.Dial("tcp", address)
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()

	att := make(map[string]string)
	for {
		servOk(conn)
		conn.Write([]byte(val + "\n"))
		m := bufio.NewReader(conn)
		for {
			line, err := m.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				} else {
					os.Exit(1)
				}
			}
			if line != "OK\n" {
				pt := strings.Split(line, ": ")
				pt[1] = strings.TrimRight(pt[1], "\n")
				att[pt[0]] = pt[1]
			} else {
				break
			}
		}
		break
	}
	return att
}
