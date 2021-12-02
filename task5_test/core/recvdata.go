package core

import (
	"log"
	"net"
)

func RecvDataHandle(conn *net.TCPConn) error {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return err
	}
	log.Printf("收到的数据长度为:%d", n)
	log.Printf("收到的数据为:%s", string(buffer))
	return nil
}
