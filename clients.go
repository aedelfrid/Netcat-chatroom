package main

import (
	"math/rand"
	"net"
)

type client struct {
	id int8
	ip net.Addr
}

type clientList []client

func newClient(conn net.Conn) client {
	user := client{
		id: int8(rand.Int()),
		ip: conn.RemoteAddr(),
	}

	return user
}