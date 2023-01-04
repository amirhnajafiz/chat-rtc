package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

// RoomCreate will send the user to a new room.
func (h *Handler) RoomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s", uuid.New().String()))
}

func (h *Handler) Room(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		c.Status(http.StatusBadRequest)
		return nil
	}

	suuid, uuid, _ := createOrGetRoom(id)
}

// RoomWebsocket creates or gets a room.
func (h *Handler) RoomWebsocket(c *websocket.Conn) {
	id := c.Params("uuid")
	if id == "" {
		return
	}

	_, _, room := createOrGetRoom(id)
}

func createOrGetRoom(uuid string) (string, string, Room) {

}