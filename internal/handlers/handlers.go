package handlers

import (
	"io"
	"log"
	"net"

	messages "github.com/mikkoryynanen/real-time/generated/proto"
	"google.golang.org/protobuf/proto"
)

var handlers = map[int32]func(*messages.WrapperMessage) {
	0: HandleClientInput,
}

func Init(conn net.Conn) {
	messagesChannel := make(chan messages.WrapperMessage, 1024)

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	go handleRequest(conn, messagesChannel)
	go handleMessages(messagesChannel)
}

func handleRequest(conn net.Conn, messageChannel chan messages.WrapperMessage) {
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Print(err)
			}
			return
		}

		// Resize the buffer to remove the extra values
		buffer = buffer[:n]

		msg := &messages.WrapperMessage{}
		if err := proto.Unmarshal(buffer, msg); err != nil {
			log.Println(err)
		}
		
		// log.Printf("message received %v", msg.GetClientInputRequest())
		
		messageChannel <- *msg
	}
}

func handleMessages(messageChannel chan messages.WrapperMessage) {
	for {
		for msg := range messageChannel {
			value, exists := handlers[msg.MessageType]
			if exists {
				value(&msg)
			} else {
				log.Printf("Could not find handler for message type %v", msg.MessageType)
			}
		}
	}
}

func HandleClientInput(msg *messages.WrapperMessage) {
	log.Println("Handling player input message")
					
	log.Printf("Received player input %v", msg.GetClientInputRequest())

	// fmt.Printf("Received PlayerInput message: %v\n", m)
}