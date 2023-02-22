package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func (h *Handler) RoomChat(c *fiber.Ctx) error {
	return c.Render("chat", fiber.Map{}, "layouts/main")
}

func (h *Handler) RoomChatWebsocket(c *websocket.Conn) {

}

func (h *Handler) StreamChatWebsocket(c *websocket.Conn) {

}
