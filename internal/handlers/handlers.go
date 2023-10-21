package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"

	message_types "github.com/mikkoryynanen/real-time/internal/types"
)

var handlers = map[int]func(net.Conn, message_types.BaseMessage) {
	0 : HandlePlayerInput,
}

func Init(conn net.Conn) {
	messagesChannel := make(chan message_types.BaseMessage, 1024)

	go handleRequest(conn, messagesChannel)
	go handleMessages(conn, messagesChannel)
}

func handleRequest(conn net.Conn, cn chan message_types.BaseMessage) {
	defer conn.Close()

	decoder := json.NewDecoder(conn)	

	for {
		var msg message_types.BaseMessage
		if err := decoder.Decode(&msg); err != nil {
			if err == io.EOF {
				log.Printf("Connection closed by the remote end.")
				break
			}
			log.Printf("Decoding failed %v", err)
		}

		cn <- msg
	}
}

func handleMessages(conn net.Conn, cn chan message_types.BaseMessage) {
	for {
		for msg := range cn {
			handlers[int(msg.Type)](conn, msg)
		}
	}
}

func HandlePlayerInput(conn net.Conn, msg message_types.BaseMessage) {
	log.Println("Handling player input message")
					
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Print(err)
	}

	// Resize the buffer to remove the extra values
	buffer = buffer[:n]

	var data message_types.PlayerInput
	if err := json.Unmarshal(buffer, &data); err != nil {
		fmt.Println("Error decoding PlayerInput data:", err)
		return
	}

	fmt.Printf("Received PlayerInput message: %v\n", data)
}