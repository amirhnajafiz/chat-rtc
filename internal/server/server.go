package server

import (
	"flag"
	"os"
	"time"

	"github.com/amirhnajafiz/churchill/internal/http/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/gofiber/websocket/v2"
)

var (
	addr = flag.String("address", ":"+os.Getenv("PORT"), "")
	cert = flag.String("cert", "", "")
	key  = flag.String("key", "", "")
)

func Run() error {
	flag.Parse()

	if *addr == ":" {
		*addr = ":8080"
	}

	// creating an engine
	engine := html.New("./views", ".html")

	// creating an app with following engine
	app := fiber.New(fiber.Config{Views: engine})

	// adding middlewares to app
	app.Use(logger.New())
	app.Use(cors.New())

	// creating a new handler
	h := handlers.Handler{}

	app.Get("/", h.Welcome)
	app.Get("/room/create", h.RoomCreate)
	app.Get("/room/:uuid", h.Room)
	app.Get("/room/:uuid/websocket", websocket.New(h.RoomWebsocket, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))
	app.Get("/room/:uuid/chat", h.RoomChat)
	app.Get("/room/:uuid/chat/websocket", websocket.New(h.RoomChatWebsocket))
	app.Get("/room/:uuid/viewer/websocket", websocket.New(h.RoomViewerWebsocket))
	app.Get("/stream/:ssuid", h.Stream)
	app.Get("/stream/:ssuid/websocket", websocket.New(h.StreamWebsocket, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))
	app.Get("/stream/:ssuid/chat/websocket", websocket.New(h.StreamChatWebsocket))
	app.Get("/stream/:ssuid/viewer/websocket", websocket.New(h.StreamViewerWebsocket))

	app.Static("/", "./assets")

	w.Rooms = make(map[string]*w.Room)
	w.Streams = make(map[string]*w.Room)

	go dispatchKeyFrames()

	if *cert != "" {
		return app.ListenTLS(*addr, *cert, *key)
	}

	return app.Listen(*addr)
}

func dispatchKeyFrames() {
	room.Peers.DispatchKeyFrame()
}
