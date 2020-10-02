package app_test

import (
	"testing"

	"github.com/nkarpenko/koho-transaction/conf"
	"github.com/nkarpenko/koho-transaction/parser"
	"github.com/nkarpenko/koho-transaction/transaction"
	"github.com/nkarpenko/koho-transaction/validator"
)

type test struct {
	result      bool
	config      *conf.Config
	parser      parser.Parser
	validator   validator.Validator
	transaction transaction.Transaction
}

func TestStart(t *testing.T) {

}

func TestNew(t *testing.T) {

}
