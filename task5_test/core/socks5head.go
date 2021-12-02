package core

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"strconv"
)

func GenSocks5Head(conn *net.TCPConn) {
	// 建立连接
	buffer := make([]byte, 256)
	n, err := conn.Write([]byte{0x5, 0x1, 0x0})
	if err != nil {
		log.Println("method send err:" + err.Error())
		return
	}
	if n != 3 {
		log.Println("method send error:too short")
		return
	}
	_, err = conn.Read(buffer[:2])
	if err != nil {
		log.Println("server response method error:" + err.Error())
		return
	}
	// 发送目标地址帧
	destAddrString := "localhost"
	destPortString := "8090"
	resp := []byte{0x05, 0x01, 0x00, 0x03}
	// IP地址
	resp = append(resp, byte(len([]byte(destAddrString))))
	resp = append(resp, []byte(destAddrString)...)
	// resp = append(resp, []byte{0x127, 0x0, 0x0, 0x1})
	// 端口
	dstPortBuff := bytes.NewBuffer(make([]byte, 0))
	dstPortInt, err := strconv.ParseUint(destPortString, 10, 16)
	if err != nil {
		log.Fatal(err)
	}
	binary.Write(dstPortBuff, binary.BigEndian, dstPortInt)
	dstPortBytes := dstPortBuff.Bytes() // int为8字节
	resp = append(resp, dstPortBytes[len(dstPortBytes)-2:]...)
	// 发送IP地址和端口
	_, err = conn.Write(resp)
	if err != nil {
		log.Fatal(err)
	}
	// 读取回复
	n, err = conn.Read(resp)
	if err != nil {
		log.Fatal(err)
	}
	var targetResp [10]byte
	copy(targetResp[:10], resp[:n])
	specialResp := [10]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	if targetResp != specialResp {
		log.Fatal("协议错误, 第二次协商返回出错")
	}
	log.Print("认证成功")
}
