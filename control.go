package main

import (
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

