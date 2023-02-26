package app

import (
	"github.com/amirhnajafiz/chat-rtc/internal/handler/http"
	"github.com/amirhnajafiz/chat-rtc/internal/protocol"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func New() {
	app := fiber.New()

	handler := http.Handler{
		Channel: make(chan protocol.Packet),
	}

	app.Static("/", "./public")

	app.Get("/ws", websocket.New(handler.Websocket))
	app.Get("/hlz", handler.Health)

	if err := app.Listen(":5001"); err != nil {
		panic(err)
	}
}
