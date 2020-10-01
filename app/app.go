package app

import (
	"errors"
	"fmt"

	"github.com/nkarpenko/koho-transaction/conf"
	"github.com/nkarpenko/koho-transaction/parser"
	tx "github.com/nkarpenko/koho-transaction/transaction"
)

var cache []int

// App struct containing required application vars.
type App struct {
	config    *conf.Config
	parser    parser.Parser
	validator tx.Validator
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
		res := a.validator.Validate(&tx)

		// Process the transaction.
		a.validator.ProcessTransaction(res)
	}

	// Final cache
	// fmt.Printf("New cache: %+v\n\n", transaction.Cache)
}

// New app instance.
func New(c *conf.Config) (*App, error) {

	// Confirm config exists.
	if c == nil {
		return &App{}, errors.New("config does not exist")
	}

	// Return new app instance.
	return &App{
		config:    c,
		parser:    parser.New(c),
		validator: tx.NewValidator(c),
	}, nil
}
