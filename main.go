package main

import (
	"github.com/damianwysocky/mmo_server/network"
)

const host = "127.0.0.1:9000"

func main() {
	server := network.NewServer(host)
	server.Start()
}
