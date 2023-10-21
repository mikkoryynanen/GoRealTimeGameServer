package network

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"

	arguments "github.com/mikkoryynanen/real-time/utils"
	handlers "github.com/mikkoryynanen/real-time/internal/handlers"
)

func Start(args *arguments.Arguments) {
	cert, err := tls.LoadX509KeyPair(args.CertPath, args.KeyPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	
	config := tls.Config{ Certificates: []tls.Certificate{cert}}
	
	listener, err := tls.Listen("tcp", fmt.Sprintf("%v:%v", args.Host, args.Port), &config)
	if err != nil {
		log.Fatal(err)
		return
	}
	
	fmt.Printf("Listening on port %v", args.Port)
	
	defer listener.Close()

	for {
		conn, err := listener.Accept()
        if err != nil {
            log.Printf("server: accept: %s", err)
            break
        }

        log.Printf("server: accepted from %s", conn.RemoteAddr())
		
        tlscon, ok := conn.(*tls.Conn)
        if ok {
            log.Print("Connection established")

            state := tlscon.ConnectionState()
            for _, v := range state.PeerCertificates {
                log.Print(x509.MarshalPKIXPublicKey(v.PublicKey))
            }
        }

		handlers.Init(conn)
	}
}