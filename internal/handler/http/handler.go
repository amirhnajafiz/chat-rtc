package http

import (
	"log"

	"github.com/amirhnajafiz/chat-rtc/internal/protocol"
	"github.com/gofiber/websocket/v2"
)

type Handler struct {
	Channel chan protocol.Packet
}

func (h *Handler) Websocket(c *websocket.Conn) {
	var (
		messageType int
		bytes       []byte
		err         error
	)

	go func() {
		for {
			data := <-h.Channel

			if err = c.WriteMessage(data.MessageType, data.Bytes); err != nil {
				log.Println("[FAIL] write:\n", err)

				break
			}
		}
	}()

	for {
		if messageType, bytes, err = c.ReadMessage(); err != nil {
			log.Println("[FAIL] read:\n", err)

			break
		}

		data := protocol.Packet{
			MessageType: messageType,
			Bytes:       bytes,
		}

		h.Channel <- data
	}
}
