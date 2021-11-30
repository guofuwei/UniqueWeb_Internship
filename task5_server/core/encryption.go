package core

import (
	"io"
	"net"
)

func EncodeWrite(conn *net.TCPConn, bs []byte) (n int, err error) {
	cipherText, err := AesEncrypt(bs)
	if err != nil {
		return
	}
	return conn.Write(cipherText)
}

func DecodeRead(conn *net.TCPConn, bs []byte) (n int, err error) {
	n, err = conn.Read(bs)
	if err != nil {
		return
	}
	bs, err = AesDecrypt(bs[0:n])
	if err != nil {
		return
	}
	return len(bs), nil
}

func EncodeCopy(src *net.TCPConn, dst *net.TCPConn) error {
	buffer := make([]byte, 128)
	for {
		readCount, readErr := src.Read(buffer)
		if readErr != nil {
			if readErr != io.EOF {
				return readErr
			} else {
				return nil
			}
		}
		if readCount > 0 {
			writeCount, writeErr := EncodeWrite(dst, buffer[0:readCount])
			if writeErr != nil {
				return writeErr
			}
			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}

func DecodeCopy(src *net.TCPConn, dst *net.TCPConn) error {
	buffer := make([]byte, 128)
	for {
		readCount, readErr := DecodeRead(src, buffer)
		if readErr != nil {
			if readErr != io.EOF {
				return readErr
			} else {
				return nil
			}
		}
		if readCount > 0 {
			writeCount, writeErr := dst.Write(buffer[0:readCount])
			if writeErr != nil {
				return writeErr
			}
			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}
