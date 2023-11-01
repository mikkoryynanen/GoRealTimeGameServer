// Simple client for testing the connection
// Not meant to used as an actual client

package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"time"

	messages "github.com/mikkoryynanen/real-time/generated/proto"
	"google.golang.org/protobuf/proto"
)

func main() {
    cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
    if err != nil {
        log.Fatalf("Failed to load certification: %s", err)
    }

    config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
    conn, err := tls.Dial("tcp", ":3000", &config)
    if err != nil {
        log.Fatalf("Dialing server failed: %s", err)
    }

    defer conn.Close()
    log.Println("Client connected to: ", conn.RemoteAddr())

    state := conn.ConnectionState()
    for _, v  := range state.PeerCertificates {
        fmt.Println(v.Subject)
    }

    go receiveMessages(conn)
    go sendMessages()

    select{}
}

func receiveMessages(conn *tls.Conn) {
    for {
        buffer := make([]byte, 1024)
        n, err := conn.Read(buffer)
        if err != nil {
            if err != io.EOF {
                log.Print(err)
            }
        }
        buffer = buffer[:n]

        if len(buffer) > 0 {
            msg := &messages.WrapperMessage{}
			if err := proto.Unmarshal(buffer, msg); err != nil {
				log.Println(err)
			}
			
			switch msg.GetMessageType() {
                case 0:
                    log.Printf("received connection response %v", msg)
                
            }
        }
    }
}

func sendMessages() {
    for {
        // TODO Create helper func?
        // msg := &messages.WrapperMessage{
        //     MessageType: 0,
        //     Msg: &messages.WrapperMessage_ClientInputRequest{
        //         ClientInputRequest: &messages.ClientInputRequest{
        //             Input: &messages.Vector2{
        //                             X: 1,
        //                             Y: 0.1,
        //                         },
        //         },
        //     },
        // }
        // sendMessage(conn, msg)

        time.Sleep(1 * time.Second)
    }
}

func sendMessage(conn *tls.Conn, msg *messages.WrapperMessage) {
    bytes, err := proto.Marshal(msg)
    if err != nil {
        log.Fatal(err)
        return
    }

    _, err = conn.Write(bytes)
    if err != nil {
        log.Fatal(err)
        return
    }

    log.Printf("Sent message %v", msg)
}