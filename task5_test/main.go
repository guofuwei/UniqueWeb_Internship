package main

import (
	"log"
	"net"
	"task5_test/core"
)

func main() {
	serverAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	serverSocket, err := net.DialTCP("tcp", nil, serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer serverSocket.Close()
	core.GenSocks5Head(serverSocket)
	err = core.SendDataHandle(serverSocket)
	if err != nil {
		log.Fatal(err)
	}
	err = core.RecvDataHandle(serverSocket)
	if err != nil {
		log.Fatal(err)
	}
}
