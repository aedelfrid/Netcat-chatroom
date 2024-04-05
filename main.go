package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Server struct {
	host string
	port string
	db *DB
	clients clientList
}

type Client struct {
	conn net.Conn
}

type clientList map[string]Client

func (*clientList) newClient(conn net.Conn) *Client {
	client := &Client{
		conn: conn,
	}

	return client
}

func (*clientList) closeClient() {

} 

type Config struct {
	Host string
	Port string
	db *DB
}

func New(config *Config) *Server {
	return &Server{
		host: config.Host,
		port: config.Port,
	}
}

func (server *Server) Run() {
	userChan := make(chan User, 10)
	messageChan := make(chan Message, 50)
	
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.host, server.port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		client := &Client{
			conn: conn,
		}

		append(server.clients, *client)

		go client.handleRequest()
	}
}

func (client *Client) handleRequest() {
	reader := bufio.NewReader(client.conn)

	// collect userinfo
	// create new Message instance
	// find user in db
	// if user not in db, add new user

	// store message in db
	// dial 

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			client.conn.Close()

			return
		}
		fmt.Printf("Message incoming: %s", string(message))
		client.conn.Write([]byte("Message received.\n"))
	}
}

func main() {
	d := initDB()
	defer d.Close()

	server := New(&Config{
		Host: "localhost",
		Port: "3333",
		db: d,
	})

	server.Run()
}


