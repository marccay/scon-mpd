package main

import (
//	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		os.Exit(1)
	} else {
		cmd := os.Args[1]
		switch cmd {
			case "play":
				send("play")
			case "stop":
				send("stop")
			case "pause":
				send("pause")
			case "next":
				send("next")
			case "previous":
				send("previous")
			default:
				os.Exit(1)
		}
	}

}

func send(command string) {
	conn, err := net.Dial("tcp", ":6600")
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()

	for  {
		servOk(conn)
		conn.Write([]byte(command + "\n"))
		servOk(conn)
		break
	}
}

func servOk(conn net.Conn) {
	m := bufio.NewReader(conn)
	ok,_ := m.ReadString('\n')
	if byte(ok[0]) != 79 || byte(ok[1]) != 75 {
		os.Exit(1)
	}
}

