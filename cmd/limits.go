package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Create the limira command.
var limitsCmd = &cobra.Command{
	Use:   "limits",
	Short: "Display user transaction limits.",
	Run: func(cmd *cobra.Command, args []string) {

		// Get the config file.
		c, err := getConfig(cmd)
		if err != nil {
			fmt.Printf("error getting config: %+v\n", err)
			return
		}

		// Print the users transaction limits from config.
		fmt.Println("User transaction limits:")
		fmt.Printf("\tMax of $%+v can be loaded per day.\n", c.Limits.DailyAmount)
		fmt.Printf("\tMax of $%+v can be loaded per week.\n", c.Limits.WeeklyAmount)
		fmt.Printf("\tMax of %+v loads per day.\n", c.Limits.DailyTransactions)
		return
	},
}
