package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strconv"
)


func previous(option string) {
	if option == "smart" || option == "s" {
		elapsed, _, song_id := song_time()
		if elapsed >= 10.00 {
			cmd := "seek " + strconv.FormatUint(song_id, 10) + " 0"
			send(cmd)
		} else {
			send("previous")
		}
	} else {
		send("previous")
	}
}

func crossfade(option string) {
	var opt int64
	opt, err := strconv.ParseInt(option, 10, 64)
	fmt.Println(opt)
	if err != nil {
		if option == "on" || option == "ON" {
			opt = 10
		} else if option == "off" || option == "OFF" {
			opt = 0
		} else if option == "" {
			current_att := single_att("xfade")
			if current_att == "" {
				opt = 10
			} else {
				opt = 0
			}
		} else {
			os.Exit(1)
		}
		cmd := "crossfade " + strconv.FormatInt(opt, 10)
		fmt.Println(cmd)
		send(cmd)
	} else {
		cmd := "crossfade " + strconv.FormatInt(opt, 10)
		fmt.Println("send int", cmd)
		send(cmd)
	}
}

func toggle(command string, option string) {
	var opt string
	current_att := single_att(command)

	if option == "" {
		if current_att == "on" {
			opt = "0"
		} else {
			opt = "1"
		}
	} else if option == "on" || option == "1" || option == "ON" {
		opt = "1"
	} else if option == "off" || option == "0" || option == "OFF" {
		opt = "0"
	} else {
		os.Exit(1)
	}

	cmd := command + " " + opt
	send(cmd) 

	if opt == "1" {
		fmt.Printf("%v is now ON\n", command)
	} else {
		fmt.Printf("%v is now OFF\n", command)
	}
}


func send(command string) {
        conn, err := net.Dial("tcp", ":6600")
        if err != nil {
                os.Exit(1)
        }
        defer conn.Close()

        for  {
                serv_ok(conn)
                conn.Write([]byte(command + "\n"))
                serv_ok(conn)
                break
        }
}

func serv_ok(conn net.Conn) {
        m := bufio.NewReader(conn)
        ok,_ := m.ReadString('\n')
        if byte(ok[0]) != 79 || byte(ok[1]) != 75 {
                os.Exit(1)
        }
}

