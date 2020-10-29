package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	buf := make([]byte, 1024)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn, buf)
	}
}
func handleClient(conn net.Conn, buf []byte) {
	defer conn.Close()
	msg, _ := conn.Read(buf)
	conn.Write(buf[:msg])

}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
