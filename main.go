package main

import (
	"flag"

	"github.com/amirhnajafiz/chat-rtc/internal"
)

func main() {
	var (
		AppPort = flag.Int("port", 5000, "application port")
	)

	flag.Parse()

	internal.NewApp(*AppPort)
}
