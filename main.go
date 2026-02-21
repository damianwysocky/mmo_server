package main

import (
	"fmt"
	"net"
)

const network = "udp"
const host = "127.0.0.1:9000"

func main() {
	conn := getConnection()
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Read error:", err)
			continue
		}

		message := string(buffer[:n])

		fmt.Printf("Got message from %v: %s\n", clientAddr, message)

		_, err = conn.WriteToUDP([]byte("Hello client!"), clientAddr)
		if err != nil {
			fmt.Println("Write error:", err)
		}
	}
}

func getConnection() *net.UDPConn {
	addr, err := net.ResolveUDPAddr(network, host)
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP(network, addr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Server started!")

	return conn
}
