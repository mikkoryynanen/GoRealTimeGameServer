package handlers

import (
	"crypto/tls"
	"log"
	"time"

	messages "github.com/mikkoryynanen/real-time/generated/proto"
	"github.com/mikkoryynanen/real-time/internal/logic"
	"github.com/mikkoryynanen/real-time/internal/session"
	"google.golang.org/protobuf/proto"
)

// TODO Not sure if this is the best thing to do
// We require this object because we need to know the client's connection when sending message
// We need to write to it's connection
type OutboundMessage struct {
	Bytes		[]byte
	Session		session.ClientSession
}
type Handler struct {
	handlers				map[int32]interface{}
	clientSessions 			[]session.ClientSession
	inboundMessagesChannel	chan []byte
	outboundMessagesChannel	chan OutboundMessage
	logic					logic.Logic
}

type MessageHandler interface {
	Handle(message *messages.WrapperMessage)
}

func NewHandler() *Handler {	
	return &Handler{}
}

func (h *Handler) Init() {
	h.handlers = make(map[int32]interface{})
	// TODO Figure out a way to automate this
	h.handlers[0] = &ClientInputHandler{
		handler: h,
	}

	h.inboundMessagesChannel = make(chan []byte, 1024)
	h.outboundMessagesChannel = make(chan OutboundMessage)
	h.logic = *logic.NewLogic()

	go h.handleInbound()
	go h.handleOutbound()
}

func (h *Handler) EnqueueMessage(buffer []byte) {
	h.inboundMessagesChannel <- buffer
}

func (h *Handler) AddConnection(conn *tls.Conn) {
	log.Println("Client connected")

	session := session.NewSession(conn)
	h.clientSessions = append(h.clientSessions, *session)

	// Send acknowledgement to client
	msg := &messages.WrapperMessage{
		MessageType: 0,
		Msg: &messages.WrapperMessage_ClientConnectionResponse{
			ClientConnectionResponse: &messages.ClientConnectionResponse{
				SessionId: session.Id,
			},
		},
	}

	// Verify that the message is built correctly
	outboundMsg := &OutboundMessage{
		Bytes:   buildMessage(msg),
		Session: *session,
	}

	// Ensure that the channel is properly initialized and available
	h.outboundMessagesChannel <- *outboundMsg

	log.Println("Ack sent to client")
}


type ClientInputHandler struct {
	handler		*Handler
}
// TODO We can receive any message type here
// add a some sort of check
func (h *ClientInputHandler) Handle(msg *messages.WrapperMessage) {
	log.Println("Handling player input message")

	input := msg.GetClientInputRequest()
	for _, session := range h.handler.clientSessions {
		if session.Id == msg.SessionId {
			h.handler.logic.HandleLogic(input, session.Id)
		}
	}

	log.Printf("Received player input %v", msg.GetClientInputRequest())
}

func (h *Handler) handleInbound() {
	for {
		for buffer := range h.inboundMessagesChannel {
			msg := &messages.WrapperMessage{}
			if err := proto.Unmarshal(buffer, msg); err != nil {
				log.Println(err)
			}
			
			handler, exists := h.handlers[msg.GetMessageType()]
			if exists {
				if handler, ok := handler.(MessageHandler); ok {
					handler.(MessageHandler).Handle(msg)
				} else {
					log.Printf("Could not start handler for message type %v", &msg)
				}
			} else {
				log.Printf("Could not find handler for message type %v", msg.GetMessageType())
			}
		}

		// TODO Not sure if we need this, seems to be working quite ok without it
		// time.Sleep(time.Millisecond)
	}
}

func (h *Handler) handleOutbound() {
	for {
		for outboundMessage := range h.outboundMessagesChannel {
			_, err := outboundMessage.Session.Conn.Write(outboundMessage.Bytes)
			if err != nil {
				log.Println(err)
				return
			}

			log.Printf("Sent message to client %v", outboundMessage.Bytes)
		}

		// TODO Not sure if we need this, seems to be working quite ok without it
		time.Sleep(time.Second)
	}
}

// TODO Might need to move this to somewhere else, the file is starting to bloat
func buildMessage(msg *messages.WrapperMessage) []byte {
	bytes, err := proto.Marshal(msg)
    if err != nil {
        log.Fatal(err)
        return nil
    }
	return bytes
}
