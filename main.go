package main

import (
	"fmt"
	"net"
)

type client struct {
	id int
	ip net.IP
}

type clientList []client

func newClient() {

}

func (c *clientList) pingClients() {

}

func handlePacket(conn net.Conn) {
	buf := make([]byte, 1024)
	conn.Read(buf)
	conn.Write([]byte("Message received.\n"))
	conn.Close()

	fmt.Println(string(buf))
}

func main() {
	const (
		network string = "tcp"
		port    string = ":8000"
	)

	clients := new(clientList)

	listener, err := net.Listen(network, port)
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()

	fmt.Printf("Listening on port %s", port)

	for {
		conn, _ := listener.Accept()
		if err != nil {
			panic(err)
		}
		clients.push(conn.RemoteAddr())

		go handlePacket(conn)
	}
}
