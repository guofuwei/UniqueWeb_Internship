package core

import (
	"errors"
	"log"
	"net"
)

func SendDataHandle(conn *net.TCPConn) error {
	testData := []byte("Participate i n the official 2021 Python Developers Survey  Loren Crary has joined the PSF as its Director of Resource Development  Python powers major aspects of Abridge’s ML lifecycle, including data annotation, research and experimentation, and ML model deployment to production. Participate in the official 2021 Python Developers Survey  Loren Crary has joined the PSF as its Director of Resource Development Participate i n the official 2021 Python Developers Survey  Loren Crary has joined the PSF as its Director of Resource Development  Python powers major aspects of Abridge’s ML lifecycle, including data annotation, research and experimentation, and ML model deployment to production.")
	n, err := conn.Write(testData)
	if err != nil {
		return errors.New("senddata error:" + err.Error())
	}
	log.Printf("发送的数据长度:%d", n)
	log.Printf("发送的数据为%s", string(testData))
	return nil
}
