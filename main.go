package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type Connection struct {
	net.Conn
}

type Message struct {
	userid int8
	body string
}

var (
	netType string = "tcp"
	addr string = ":8000"
	clients chan clientList
	connectedClients chan clientList
	lastMessage chan Message
)

func (m Message)sendMessage(sender client) {
	for _, client := range <-connectedClients {
		if client.id != m.userid{
			return
		}

		d, _ := net.Dial(client.ip.Network(), client.ip.Network())
		d.Write([]byte(m.body))
		
	}
}

func pingClients() {
	mc := <-clients
	buf := make([]byte, 64)

	for i, c := range mc {
		d, err := net.Dial(c.ip.Network(), c.ip.String())
		if err != nil {
			log.Print(err)
		}

		d.Write([]byte("ping"))
		d.Read(buf)
		d.Close()

		if string(buf) != "pong" {
			mc = append(mc[:i], mc[i+1:]...)
		}
	}

	connectedClients <- mc
}

func handleReq(conn net.Conn) {
	cl := <-clients

	sender := newClient(conn)

	cl = append(cl, sender)
	clients <- cl

	reader := bufio.NewReader(conn)
	for {
		m, err := reader.ReadString('\n')
		if err != nil {
			conn.Close()
			return
		}

		message := Message{
			userid: sender.id,
			body: m,
		}

		lastMessage <- message
	}
}

func server() {
	l, err := net.Listen(netType, addr)
	if err != nil {
		panic(err)
	}
	defer l.Close()


	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
		}

		go handleReq(conn)
	}
}

func main() {
	tick := time.NewTicker(60 * time.Second)

	for range tick.C {
		go pingClients()
	}

	
}
