package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

type Client struct {
	id int
	ip net.Addr
}

type ClientList []Client

func (clients ClientList) newClient(conn net.Conn) Client {
	
	client := Client{
		ip: conn.RemoteAddr(),
	}

	clients = append(clients, client)

	return client
}

func pingClients(clients ClientList) ClientList{
	for i, client := range clients {
		conn, err := net.Dial(client.ip.Network(), client.ip.String())
		if err != nil {
			fmt.Println(err)
		}
		buf := make([]byte, 64)
		

		conn.Write([]byte("Ping"))
		conn.Read(buf)
		conn.Close()

		if string(buf) != "Pong" {
			clients = append(clients[:i], clients[i+1] )
		}
	}

	return clients
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

	clients := make(ClientList, 42)

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
		
		user := clients.newClient(conn)

		go handlePacket(conn)
	}
}
