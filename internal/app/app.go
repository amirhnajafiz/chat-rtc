package app

import (
	"github.com/amirhnajafiz/chat-rtc/internal/handler/http"
	"github.com/amirhnajafiz/chat-rtc/internal/handler/socket"
	"github.com/amirhnajafiz/chat-rtc/internal/protocol"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"gopkg.in/olahol/melody.v1"
)

func New() {
	app := fiber.New()
	r := gin.Default()

	ws := socket.WebSocket{
		Melody: melody.New(),
	}
	ws.Melody.HandleMessage(ws.MessageHandler)

	h := http.Handler{
		WS: &ws,
	}

	v2h := http.HandlerV2{
		Channel: make(chan protocol.Packet),
	}

	r.Use(static.Serve("/", static.LocalFile("./public", true)))
	r.GET("/ws", h.OpenSocket)
	r.GET("/healthz", h.Health)

	app.Get("/ws", websocket.New(v2h.Websocket))

	go func() {
		if err := app.Listen(":5001"); err != nil {
			panic(err)
		}
	}()

	if err := r.Run(":5000"); err != nil {
		panic(err)
	}
}
