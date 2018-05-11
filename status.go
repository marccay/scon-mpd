package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"
	"net"
	"strconv"
)


func single_att(attribute string) string {
        att := status()
	if attribute == "state" || attribute == "duration" || attribute == "elapsed" {
		return att[attribute]
	} else if attribute == "time" {
		return att["elapsed"] + "/" + att["duration"]
	} else {
		is_On, _ := is_att_on(att[attribute])
		return is_On
	}
}

func is_att_on(stat string) (string, int) {
	if stat == "0" {
		return "off", 0
	} else {
		return "on", 1
	}
}

func song_time() (float64, float64, uint64) {
	att := status()
	duration, err := strconv.ParseFloat(att["duration"], 64)
	if err != nil {
		os.Exit(1)
	}
	elapsed, err := strconv.ParseFloat(att["elapsed"], 64)
	if err != nil {
		os.Exit(1)
	}
	song_id, err := strconv.ParseUint(att["song"], 10, 64) 
	if err != nil {
		os.Exit(1)
	}
	return elapsed, duration, song_id 

}


func print_status() {
        att := status()
	fmt.Println("state:", att["state"])
	is_On, _ := is_att_on(att["repeat"])
	fmt.Println("repeat:", is_On)
	is_On, _ = is_att_on(att["random"])
	fmt.Println("random:", is_On)
	is_On, _ = is_att_on(att["single"])
	fmt.Println("single:", is_On)
	is_On, _ = is_att_on(att["consume"])
	fmt.Println("consume:", is_On)
	is_On, _ = is_att_on(att["xfade"])
	fmt.Println("crossfade:", is_On)

	if att["state"] == "play" || att["state"] == "pause" {
		fmt.Println("elapsed", att["elapsed"])
		fmt.Println("duration", att["duration"])
	}
}

func status() map[string]string {
        conn, err := net.Dial("tcp", ":6600")
        if err != nil {
                os.Exit(1)
        }
        defer conn.Close()

        att := make(map[string]string)
        for {
                serv_ok(conn)
                conn.Write([]byte("status\n"))
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
                                pt  := strings.Split(line, ": ")
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

