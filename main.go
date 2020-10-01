package main

import (
	"log"

	"github.com/nkarpenko/koho-transaction/cmd"
)

func main() {
	// Execute main root cobra command.
	if err := cmd.RootCmd().Execute(); err != nil {
		log.Printf("Unexpected error: %s\n", err.Error())
		return
	}
}
