package http

import (
	"github.com/amirhnajafiz/chat-rtc/internal/handler/socket"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	WS *socket.WebSocket
}

func (h *Handler) Health(c *gin.Context) {
	_ = c.String("OK")	
}

func (h *Handler) OpenSocket(c *gin.Context) {
	if err := h.WS.Melody.HandleRequest(c.Writer, c.Request); err != nil {
		_ = c.Error(err)
	}
}
