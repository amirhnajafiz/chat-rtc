package chat

import (
	"github.com/gofiber/websocket/v2"
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}
