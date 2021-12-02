package core

import (
	"errors"
	"log"
	"net"
)

func SendDataHandle(conn *net.TCPConn) error {
	testData := []byte("this is a test data")
	n, err := conn.Write(testData)
	if err != nil {
		return errors.New("senddata error:" + err.Error())
	}
	log.Printf("发送的数据长度:%d", n)
	log.Printf("发送的数据为%s", string(testData))
	return nil
}
