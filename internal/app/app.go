package app

import (
	"github.com/amirhnajafiz/chat-rtc/internal/handler/http"
	"github.com/amirhnajafiz/chat-rtc/internal/handler/socket"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func New() {
	r := gin.Default()

	ws := socket.WebSocket{
		Melody: melody.New(),
	}
	ws.Melody.HandleMessage(ws.MessageHandler)

	h := http.Handler{
		WS: &ws,
	}

	r.Use(static.Serve("/", static.LocalFile("./public", true)))
	r.GET("/ws", h.OpenSocket)
	r.GET("/healthz", h.Health)

	if err := r.Run(":5000"); err != nil {
		panic(err)
	}
}
