package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	//_, err = conn.Write([]byte(conn.LocalAddr().String() + " join the chatroom"))
	checkError(err)
	go receiveMsg(conn)
	for {
		var data string
		fmt.Scan(&data)
		if data == "quit" {
			break
		}
		sendmsg := []byte(data + "\n")
		conn.Write(sendmsg)
	}

}
func receiveMsg(conn *net.TCPConn) {

	result := make([]byte, 256)
	for {
		_, err := conn.Read(result)
		checkError(err)
		fmt.Print(string(result))
	}

	/*
		reader := bufio.NewReader(conn)
		for {
			msg, err := reader.ReadString('\n')
			fmt.Print(msg)
			checkError(err)
		}
	*/
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
