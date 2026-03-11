package network

import (
	"fmt"
	"net"
	"sync"
)

const protocol = "udp"

type Server struct {
	conn       *net.UDPConn
	bufferPool sync.Pool
}

func NewServer(host string) *Server {
	addr, err := net.ResolveUDPAddr(protocol, host)
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP(protocol, addr)
	if err != nil {
		panic(err)
	}

	server := &Server{
		conn: conn,
		bufferPool: sync.Pool{
			New: func() any {
				return make([]byte, 1024)
			},
		},
	}

	fmt.Println("Server started!")

	return server
}

func (s *Server) Start() {
	for {
		buffer := s.bufferPool.Get().([]byte)

		n, clientAddr, err := s.conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Read error:", err)
			s.bufferPool.Put(buffer)
			continue
		}

		message := string(buffer[:n])
		fmt.Printf("Got message from %v: %s\n", clientAddr, message)

		_, err = s.conn.WriteToUDP([]byte("Hello client!"), clientAddr)
		if err != nil {
			fmt.Println("Write error:", err)
		}

		s.bufferPool.Put(buffer)
	}
}
