package core

import (
	"io"
	"log"
	"net"
)

func EncodeWrite(conn *net.TCPConn, bs []byte) (n int, err error) {
	cipherText, err := Encrypt(bs)
	if err != nil {
		return
	}
	// log.Println("client send:")
	// log.Println(cipherText)
	return conn.Write(cipherText)
}

func DecodeRead(conn *net.TCPConn, bs []byte) (plainText []byte, n int, err error) {
	n, err = conn.Read(bs)
	if err != nil {
		return
	}
	plainText, err = Decrypt(bs[0:n])
	if err != nil {
		return
	}
	return plainText, len(plainText), nil
}

func EncodeCopy(src *net.TCPConn, dst *net.TCPConn) error {
	buffer := make([]byte, 1024)
	for {
		readCount, readErr := src.Read(buffer)
		// log.Println(buffer[:10])
		if readErr != nil {
			if readErr != io.EOF {
				return readErr
			} else {
				return nil
			}
		}
		if readCount > 0 {
			_, writeErr := EncodeWrite(dst, buffer[0:readCount])
			if writeErr != nil {
				return writeErr
			}
			// if readCount != writeCount {
			// 	log.Printf("EncodeCopy:readCount:%d\n", readCount)
			// 	log.Printf("EncodeCopy:writecount:%d\n", writeCount)
			// 	return io.ErrShortWrite
			// }
		}
	}
}

func DecodeCopy(src *net.TCPConn, dst *net.TCPConn) error {
	buffer := make([]byte, 1024)
	for {
		plainText, readCount, readErr := DecodeRead(src, buffer)
		// log.Println("client recv:")
		// log.Println(plainText)
		if readErr != nil {
			if readErr != io.EOF {
				return readErr
			} else {
				return nil
			}
		}
		if readCount > 0 {
			writeCount, writeErr := dst.Write(plainText)
			if writeErr != nil {
				return writeErr
			}
			if readCount != writeCount {
				log.Printf("DecodeCopy:readCount:%d\n", readCount)
				log.Printf("DecodeCopy:writecount:%d\n", writeCount)
				return io.ErrShortWrite
			}
		}
	}
}
