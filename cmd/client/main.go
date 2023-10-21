// Simple client for testing the connection
// Not meant to used as an actual client

package main

import (
	"math/rand"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"time"

	message_types "github.com/mikkoryynanen/real-time/internal/types"
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
        msg := &message_types.PlayerInput{
            X: int16(rand.Intn(101)),
            Y: int16(rand.Intn(101)),
        }

        json, err := json.Marshal(msg)
        if err != nil {
            log.Fatal(err)
            return
        }

        _, err = conn.Write(json)
        if err != nil {
            log.Fatal(err)
            return
        }

        time.Sleep(1 * time.Second)
    }
}