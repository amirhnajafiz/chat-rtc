package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// HTTP manages the execution of http server for Churchill.
type HTTP struct{}

// Command returns the cobra command of HTTP.
func (h HTTP) Command() *cobra.Command {
	run := func(cmd *cobra.Command, args []string) { h.main() }

	return &cobra.Command{Use: "execute http server", Run: run}
}

// main is the main function of HTTP.
func (h HTTP) main() {
	log.Println("http server started")
}
