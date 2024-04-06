package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
	"time"
)

type Server struct {
	host        string
	port        string
	db          *DB
	clients     clientList
	clientChan  chan Client
	messageChan chan Message
}

type Message struct {
	Username  string
	TimeStamp time.Time
	Body      string
}

type Client struct {
	conn net.Conn
	server *Server
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
	db   *DB
}

func New(config *Config) *Server {
	return &Server{
		host: config.Host,
		port: config.Port,
	}
}

func (server *Server) Run() {
	server.clientChan = make(chan Client, 10)
	server.messageChan = make(chan Message, 50)
	server.clients = make(clientList)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.host, server.port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			conn.Close()
			return
		}

		client := &Client{
			conn: conn,
			server: server,
		}

		go client.handleRequest(server)
	}
}

func (client *Client) handleRequest(s *Server) {
	reader := bufio.NewReader(client.conn)

	// collect userinfo
	// create new Message instance
	// find user in db
	// if user not in db, add new user

	// store message in db
	// dial

	for {
		jsonData, err := reader.ReadString('\n')
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("Message incoming: %s", jsonData)
		client.conn.Write([]byte("Message received.\n"))

		message := &Message{}

		_ = json.Unmarshal([]byte(jsonData), &message)

		// append(server.clients, *client)
		s.clients[message.Username] = *client
	}
}

func main() {
	tick := time.NewTicker(60 * time.Second)

	server := New(&Config{
		Host: "localhost",
		Port: "3333",
		db:   d,
	})

	server.Run()
}
