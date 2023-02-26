package internal

import (
	"fmt"

	"github.com/amirhnajafiz/chat-rtc/internal/handler/http"
	"github.com/amirhnajafiz/chat-rtc/internal/protocol"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func NewApp(port int) {
	app := fiber.New()

	handler := http.Handler{
		Channel: make(chan protocol.Packet),
	}

	app.Static("/", "./public")

	app.Get("/hlz", handler.Health)
	app.Get("/ws", websocket.New(handler.Websocket))

	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}
