// Package parser is a collection of file related parsing interfaces and
// methods to assist in handling user transactions.
package parser

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/nkarpenko/koho-transaction/conf"
	txmodel "github.com/nkarpenko/koho-transaction/transaction/model"
)

// Parser interface contains methods required to parsing input files.
type Parser interface {
	ParseFile() (*[]txmodel.Transaction, error)
}

type parser struct {
	input string
}

// ParseFile method parses files with a collection of JSON objects
// and returns a slice of transaction structs.
func (p *parser) ParseFile() (*[]txmodel.Transaction, error) {
	var txs []txmodel.Transaction

	// Confirm input file path exists in config.
	if p.input == "" {
		return nil, errors.New("invalid input file path supplied in config")
	}

	// Try and open the input file.
	file, err := os.Open(p.input)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)

	// Loop through file line by line to
	var line string
	for {

		// Get the new line.
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		// Unmarshal string to transaction model.
		tx := &txmodel.Transaction{}
		err := json.Unmarshal([]byte(line), tx)
		if err != nil {
			break
		}

		// Add transaction to the main list of transactions
		txs = append(txs, *tx)
	}

	// If error was end of file then it was expected.
	if err != io.EOF {
		return nil, err
	}

	// Successful input file parse.
	return &txs, nil
}

// New parser instance initialization.
func New(c *conf.Config) Parser {
	return &parser{
		input: c.InputFile,
	}
}
