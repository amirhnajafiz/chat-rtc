package handlers

import (
	"fmt"
	"net/http"

	w "github.com/amirhnajafiz/churchill/pkg/webrtc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

// RoomCreate will send the user to a new room.
func (h *Handler) RoomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s", uuid.New().String()))
}

// Room creates or gets a new room by getting room id from user request.
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
	w.RoomConn(c, room.Peers)
}

// createOrGetRoom generates a new room.
func createOrGetRoom(uuid string) (string, string, w.Room) {

}

// RoomViewerWebsocket returns room websocket.
func (h *Handler) RoomViewerWebsocket(c *websocket.Conn) {

}

func roomViewerConn(c *websocket.Conn, p *w.Peers) {

}

type websocketMessage struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}
