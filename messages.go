package main

import (
	"encoding/json"
	"os"
	"time"
)

type Message struct {
	userid uint8
	timeStamp time.Time
	body string
}

func (message *Message)messageToJSON() {
	file, _ := json.MarshalIndent(message, "", " ")

	_ = os.WriteFile("messages.json", file, 0644)
}