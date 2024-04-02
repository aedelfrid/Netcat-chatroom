package main

import (
	"math/rand"
	"net"
)

type User struct {
	id int8
	ip net.Addr
}

type userList []User

func newUser(conn net.Conn) User {
	user := User{
		id: int8(rand.Int()),
		ip: conn.RemoteAddr(),
	}

	return user
}