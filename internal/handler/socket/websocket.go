package socket

import (
	"log"

	"gopkg.in/olahol/melody.v1"
)

type WebSocket struct {
	Melody *melody.Melody
}

func (w *WebSocket) MessageHandler(session *melody.Session, bytes []byte) {
	if err := w.Melody.Broadcast(bytes); err != nil {
		log.Printf("broadcast failed:\n%s\n%v", session.Request.RemoteAddr, err)
	}
}
