package main

import (
	"log"
	
	"github.com/amirhnajafiz/chat-rtc/internal/app"
)

func main() {
	log.Println("starting app on port 5000 ...")
	
	app.New()
}
