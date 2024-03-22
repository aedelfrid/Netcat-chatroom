package main

import (
	"fmt"
	"net"
)

func main() {
	const (
		network string = "TCP"
		port    string = "8080"
	)

	packet, err := net.Listen(network, port)

	if (err != nil) {
		fmt.Println(err)
	}

	
}
