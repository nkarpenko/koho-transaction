package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Create the version command.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display app version.",
	Run: func(cmd *cobra.Command, args []string) {

		// Get the config file.
		c, err := getConfig(cmd)
		if err != nil {
			fmt.Printf("error getting config: %+v\n", err)
			return
		}

		fmt.Println("Koho user transaction tool v" + c.Version)
		return
	},
}
