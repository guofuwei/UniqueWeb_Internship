package main

import (
	"log"
	"net"
)

func main() {
	listenAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatal(err)
	}
	listenSocket, err := net.ListenTCP("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer listenSocket.Close()
	// 先打印原数据
	// testData := []byte("this is a test data")
	// log.Printf("原数据长度为:%d\n", len(testData))
	// log.Printf("原数据为:%s\n", string(testData))
	for {
		clientSocket, err := listenSocket.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		go HandleClient(clientSocket)
	}
}

func HandleClient(client *net.TCPConn) {
	defer client.Close()
	buffer := make([]byte, 1024*1024)
	n, err := client.Read(buffer)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("接受的数据长度为:%d\n", n)
	log.Printf("接受的数据为:%s", string(buffer))
	recvData := []byte("I have recv the data")
	n, err = client.Write(recvData)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("发送的数据长度为%d", n)
	log.Printf("发送的数据为%s", string(recvData))
}
