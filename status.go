package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strings"
	"net"
)



func single_att(attribute string) string {
        att := status()
        return att[attribute]
}

func print_status() {
        attributes := status()
        for i, x := range attributes {
                fmt.Println(i,x)
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

