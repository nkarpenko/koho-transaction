package main

import (
	"fmt"

	"github.com/nkarpenko/koho-transaction/cmd"
)

func main() {
	// Execute main root cobra command.
	if err := cmd.RootCmd().Execute(); err != nil {
		fmt.Printf("failed to initialize application: %s\n", err.Error())
		return
	}
}
