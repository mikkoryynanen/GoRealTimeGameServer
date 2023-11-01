package network

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"

	handlers "github.com/mikkoryynanen/real-time/internal/handlers"
	arguments "github.com/mikkoryynanen/real-time/utils"
)

func Start(args *arguments.Arguments) error {
	cert, err := tls.LoadX509KeyPair(args.CertPath, args.KeyPath)
	if err != nil {
		return err
	}
	
	config := tls.Config{ Certificates: []tls.Certificate{cert}}
	
	listener, err := tls.Listen("tcp", fmt.Sprintf("%v:%v", args.Host, args.Port), &config)
	if err != nil {
		return err
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
			
			go readStream(tlsconn, handler)

            // state := tlscon.ConnectionState()
            // for _, v := range state.PeerCertificates {
            //     log.Print(x509.MarshalPKIXPublicKey(v.PublicKey))
            // }
        }		
	}

	return nil
}

func readStream(conn *tls.Conn, handler *handlers.Handler) {
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Print(err)
			}
			handler.RemoveClient(conn)
		}

		buffer = buffer[:n]

		if len(buffer) > 0 {
			handler.EnqueueMessage(buffer)
			}
		}
}