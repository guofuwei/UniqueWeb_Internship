package main

import (
	"log"
	"test/aes"
)

// func EncodeWrite(conn *net.TCPConn, bs []byte) (n int, err error) {
// 	cipherText, err := aes.Encrypt(bs)
// 	if err != nil {
// 		return
// 	}
// 	return conn.Write(cipherText)
// }

// func DecodeRead(conn *net.TCPConn, bs []byte) (n int, err error) {
// 	n, err = conn.Read(bs)
// 	if err != nil {
// 		return
// 	}
// 	bs, err = aes.Decrypt(bs[0:n])
// 	if err != nil {
// 		return
// 	}
// 	return len(bs), nil
// }

// func EncodeCopy(src *net.TCPConn, dst *net.TCPConn) error {
// 	buffer := make([]byte, 128)
// 	for {
// 		readCount, readErr := src.Read(buffer)
// 		if readErr != nil {
// 			if readErr != io.EOF {
// 				return readErr
// 			} else {
// 				return nil
// 			}
// 		}
// 		if readCount > 0 {
// 			writeCount, writeErr := EncodeWrite(dst, buffer[0:readCount])
// 			if writeErr != nil {
// 				return writeErr
// 			}
// 			if readCount != writeCount {
// 				return io.ErrShortWrite
// 			}
// 		}
// 	}
// }

// func DecodeCopy(src *net.TCPConn, dst *net.TCPConn) error {
// 	buffer := make([]byte, 128)
// 	for {
// 		readCount, readErr := DecodeRead(src, buffer)
// 		if readErr != nil {
// 			if readErr != io.EOF {
// 				return readErr
// 			} else {
// 				return nil
// 			}
// 		}
// 		if readCount > 0 {
// 			writeCount, writeErr := dst.Write(buffer[0:readCount])
// 			if writeErr != nil {
// 				return writeErr
// 			}
// 			if readCount != writeCount {
// 				return io.ErrShortWrite
// 			}
// 		}
// 	}
// }

func main() {
	testData := []byte("this is a test data")
	log.Println(testData)
	println("-----------------------------------")
	testData, err := aes.Encrypt(testData)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(testData)
	println("----------------------------------")
	testData, err = aes.Decrypt(testData)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(testData)
	log.Println(string(testData))
}
