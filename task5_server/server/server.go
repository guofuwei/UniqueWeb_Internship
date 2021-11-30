package server

import (
	"errors"
	"log"
	"net"
	"task5_server/config"
	"task5_server/core"
)

func ListenServer(listenAddr *net.TCPAddr) {
	serverListener, err := net.ListenTCP("tcp", listenAddr)
	if err != nil {
		log.Fatal("服务器监听端口错误")
	}
	defer serverListener.Close()
	for {
		serverClient, err := serverListener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		serverClient.SetLinger(0)
		go handleServerClient(serverClient)
	}
}

func handleServerClient(serverClient *net.TCPConn) {
	err := localAuthHandle(serverClient)
	if err != nil {
		log.Println(err)
		serverClient.Close()
		return
	}
	destSocket, err := localDestHandle(serverClient)
	if err != nil {
		log.Println(err)
		serverClient.Close()
		return
	}
	if err != nil {
		log.Println(err)
		serverClient.Close()
		return
	}
	// 双向转发
	go func() {
		err = core.EncodeCopy(destSocket, serverClient)
		if err != nil {
			destSocket.Close()
			serverClient.Close()
			return
		}
	}()
	go func() {
		err = core.DecodeCopy(serverClient, destSocket)
		if err != nil {
			destSocket.Close()
			serverClient.Close()
			return
		}
	}()
}

func localAuthHandle(serverClient *net.TCPConn) error {
	// 开始进行身份验证
	// 读取uLen,uName,pLen,passwd
	buffer := make([]byte, 128)
	_, err := serverClient.Read(buffer[:1])
	if err != nil {
		return err
	}
	uLen := int(buffer[0])
	n, err := serverClient.Read(buffer[:uLen])
	if err != nil {
		return err
	}
	if n != uLen {
		return errors.New("read uname error")
	}
	uName := string(buffer[:uLen])

	_, err = serverClient.Read(buffer[:1])
	if err != nil {
		return err
	}
	pLen := int(buffer[0])
	n, err = serverClient.Read(buffer[:pLen])
	if err != nil {
		return err
	}
	if n != pLen {
		return errors.New("read passwd error")
	}
	passwd := string(buffer[:pLen])

	// 读取uName和passwd完成，开始比对uName和passwd并回复客户端
	if uName == config.USERNAME && passwd == config.PASSWORD {
		serverClient.Write([]byte{0x0})
	} else {
		serverClient.Write([]byte{0x1})
	}
	return nil
}

func localDestHandle(serverClient *net.TCPConn) (*net.TCPConn, error) {
	// 读取destLen,destAddr
	buffer := make([]byte, 128)
	_, err := serverClient.Read(buffer[:1])
	if err != nil {
		return nil, errors.New("get destLen error :" + err.Error())
	}
	destLen := int(buffer[0])
	n, err := serverClient.Read(buffer[:destLen])
	if err != nil {
		return nil, errors.New("read destAddr error :" + err.Error())
	}
	if n != destLen {
		return nil, errors.New("read destAddr length error")
	}
	destAddrString := string(buffer[:destLen])
	destAddr, err := net.ResolveTCPAddr("tcp", destAddrString)
	if err != nil {
		return nil, err
	}
	destSocket, err := net.DialTCP("tcp", nil, destAddr)
	if err != nil {
		return nil, err
	}
	return destSocket, nil
}
