package session

import (
	"crypto/tls"

	"github.com/google/uuid"
)

type ClientSession struct {
	Id			string
	Conn		*tls.Conn
}

func NewSession(conn *tls.Conn) *ClientSession {
	return &ClientSession{
		Id: uuid.NewString(),
		Conn: conn,
	}
}
