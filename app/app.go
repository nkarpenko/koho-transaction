// Package app holds a collection configurations, interfaces and methods
// required to run the app/tool.
package app

import (
	"errors"
	"fmt"

	"github.com/nkarpenko/koho-transaction/conf"
	"github.com/nkarpenko/koho-transaction/parser"
	"github.com/nkarpenko/koho-transaction/transaction"
	"github.com/nkarpenko/koho-transaction/validator"
)

var cache []int

// App struct containing required application vars.
type App struct {
	config      *conf.Config
	parser      parser.Parser
	validator   validator.Validator
	transaction transaction.Transaction
}

// Start the application.
func (a *App) Start() {

	// Parse input file to retrieve all the .
	txs, err := a.parser.ParseFile()
	if err != nil {
		fmt.Printf("Failed to parse file. Error: %v\n", err)
		return
	}

	// Loop through each transaction and try to validate + process it.
	for _, tx := range *txs {

		// Validate the transaction.
		res := a.transaction.Validate(&tx)

		// Process the transaction.
		a.transaction.Process(res)
	}
}

// New app instance.
func New(c *conf.Config) (*App, error) {

	// Confirm config exists.
	if c == nil {
		return &App{}, errors.New("config does not exist")
	}

	// Return new app instance.
	return &App{
		config:      c,
		parser:      parser.New(c),
		validator:   validator.New(c),
		transaction: transaction.New(c),
	}, nil
}
