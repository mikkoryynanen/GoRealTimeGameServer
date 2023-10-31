// Simple client for testing the connection
// Not meant to used as an actual client

package main

import (
	"crypto/tls"
	"fmt"
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
    
    for {
        // TODO Create helper func?
        msg := &messages.WrapperMessage{
            MessageType: 0,
            Msg: &messages.WrapperMessage_ClientInputRequest{
                ClientInputRequest: &messages.ClientInputRequest{
                    Input: &messages.Vector2{
                                    X: 1,
                                    Y: 0.1,
                                },
                },
            },
        }
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

        log.Println("Sent message %v", msg)

        time.Sleep(1 * time.Second)
    }
}