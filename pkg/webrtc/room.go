package webrtc

import (
	"log"
	"sync"

	"github.com/amirhnajafiz/churchill/pkg/chat"
	"github.com/gofiber/websocket/v2"
)

type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}

func RoomConn(c *websocket.Conn, p *Peers) {
	var config webrtc.Configuration

	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Println(err)

		return
	}

	newPeer := PeerConnectionState{
		PeerConnection: peerConnection,
		WebSocket:      &ThreadSafeWriter{},
		Conn:           c,
		Mutex:          sync.Mutex{},
	}
}