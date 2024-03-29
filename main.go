package main

import (
	"fmt"
	"net"
	"time"
)

type Client struct {
	id int
	ip net.Addr
}

type ClientList []Client

func (clients ClientList) newClient(conn net.Conn) (Client) {
	
	client := Client{
		ip: conn.RemoteAddr(),
	}

	return client
}

func (clients ClientList) pingClients() ClientList{
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

func handlePayload(conn net.Conn, sender Client) {
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

	tick := time.NewTicker(60 * time.Second)
	clients := make(ClientList, 42)

	for range tick.C {
		clients = clients.pingClients()
	}

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
		
		sender := clients.newClient(conn)
		clients = append(clients, sender)

		go handlePayload(conn, sender)
	}
}
