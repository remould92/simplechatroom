package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

var connSlice []net.Conn

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		connSlice = append(connSlice, conn)

		go handleClient(conn)
	}
}
func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	request := make([]byte, 128)
	defer conn.Close()

	for {
		read_len, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		if read_len == 0 || conn == nil {
			for i := 0; i < len(connSlice); i++ {
				if connSlice[i] == conn {
					connSlice = append(connSlice[:i], connSlice[i+1:]...)
					break
				}
			}

		} else {
			for i := 0; i < len(connSlice); i++ {
				if connSlice[i] == conn {
					continue
				} else {
					_, err := connSlice[i].Write([]byte(conn.RemoteAddr().String() + ":  " + string(request[:read_len])))
					checkError(err)
				}
			}
		}
		request = make([]byte, 128)
	}

}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
