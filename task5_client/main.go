package main

import (
	"flag"
	"task5_client/client"
)

func main() {
	listenAddr := flag.String("local", ":8080", "Input a port to listen:")
	serverAddr := flag.String("server", "110.40.178.201:8080", "Input remote server address:")
	agreeMent := flag.String("agree", "socks5", "Input a agreement:")
	flag.Parse()
	client.ListenLocal(*listenAddr, *serverAddr, *agreeMent)
}
