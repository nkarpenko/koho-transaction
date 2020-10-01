package cmd

import (
	"fmt"

	"github.com/nkarpenko/koho-transaction/app"

	"github.com/nkarpenko/koho-transaction/conf"
	"github.com/spf13/cobra"
)

// RootCmd execution.
func RootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "koho-transaction",
		Short: "Koho transaction validation tool.",
		Run: func(cmd *cobra.Command, args []string) {

			// Get the config file.
			c, err := getConfig(cmd)
			if err != nil {
				fmt.Printf("error getting config: %+v\n", err)
				return
			}

			// Start the app.
			start(c)
		},
	}

	// Add any additional flags.
	rootCmd.PersistentFlags().StringP("config", "c", "config.yml", "Specify local configuration file.")

	// Add additional commands.
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(limitsCmd)

	return rootCmd
}

// Start the application.
func start(c *conf.Config) {

	// Init the app.
	a, err := app.New(c)
	if err != nil {
		fmt.Printf("App failed to start. Error: %v\n", err)
		return
	}

	// Start the app.
	a.Start()
}

func getConfig(cmd *cobra.Command) (*conf.Config, error) {

	// Get the config file.
	configFile, err := cmd.Flags().GetString("config")
	if err != nil {
		fmt.Printf("invalid CLI flags, please use the -h flag to see all available options: %+v\n", err)
		return &conf.Config{}, err
	}

	// Load the config file.
	config, err := conf.Load(configFile)
	if err != nil {
		fmt.Printf("failed to load configration file: %+v\n", err)
		return &conf.Config{}, err
	}

	// Successful config request.
	return config, nil
}
