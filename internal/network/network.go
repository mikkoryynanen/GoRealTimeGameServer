package network

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"

	handlers "github.com/mikkoryynanen/real-time/internal/handlers"
	arguments "github.com/mikkoryynanen/real-time/utils"
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

	handler := handlers.NewHandler()
	handler.Init()

	// Continously accept connections
	for {
		conn, err := listener.Accept()
        if err != nil {
            log.Printf("server: accept: %s", err)
            break
        }

        log.Printf("Connection accepted from %s", conn.RemoteAddr())
        
        if tlsconn, ok := conn.(*tls.Conn); ok {
			handler.AddConnection(tlsconn)

			defer tlsconn.Close()
			
			for {
				buffer := make([]byte, 1024)
				n, err := tlsconn.Read(buffer)
				if err != nil {
					if err != io.EOF {
						log.Print(err)
					}
					return
				}

				buffer = buffer[:n]

				handler.EnqueueMessage(buffer)
			}

            // state := tlscon.ConnectionState()
            // for _, v := range state.PeerCertificates {
            //     log.Print(x509.MarshalPKIXPublicKey(v.PublicKey))
            // }
        }		
	}
}