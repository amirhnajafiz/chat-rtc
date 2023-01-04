package main

import (
	"github.com/amirhnajafiz/churchill/cmd"
	"github.com/spf13/cobra"
)

// main function of Churchill application.
// in this function we start the application using cobra.
func main() {
	// create the root command
	root := &cobra.Command{}

	// add cmd commands
	root.AddCommand(
		cmd.HTTP{}.Command(),
	)

	// execute root command
	if err := root.Execute(); err != nil {
		panic(err)
	}
}
