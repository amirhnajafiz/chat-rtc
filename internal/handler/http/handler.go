package http

import (
	"log"
	"net/http"

	"github.com/amirhnajafiz/chat-rtc/internal/handler/socket"
	"github.com/amirhnajafiz/chat-rtc/internal/protocol"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/websocket/v2"
)

type Handler struct {
	WS *socket.WebSocket
}

func (h *Handler) Health(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func (h *Handler) OpenSocket(c *gin.Context) {
	if err := h.WS.Melody.HandleRequest(c.Writer, c.Request); err != nil {
		_ = c.Error(err)
	}
}

type HandlerV2 struct {
	Channel chan protocol.Packet
}

func (h *HandlerV2) Websocket(c *websocket.Conn) {
	var (
		messageType int
		bytes       []byte
		err         error
	)

	for {
		if messageType, bytes, err = c.ReadMessage(); err != nil {
			log.Println("[FAIL] read:\n", err)

			continue
		}

		if err = c.WriteMessage(messageType, bytes); err != nil {
			log.Println("[FAIL] write:\n", err)

			continue
		}
	}
}
