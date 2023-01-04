package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket"
	"github.com/google/uuid"
)

// RoomCreate will send the user to a new room.
func (h *Handler) RoomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s", uuid.New().String()))
}

// RoomWebsocket .
func (h *Handler) RoomWebsocket(c *websocket.Conn) {
	id := c.Params("uuid")
	if id == "" {
		return
	}

	createOrGetRoom(id)
}