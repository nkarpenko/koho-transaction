// Package transaction contains the main service for the KOHO transaction tool
// along with a collection of related methods and helper functions.
package transaction

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/nkarpenko/koho-transaction/common/cache"
	"github.com/nkarpenko/koho-transaction/common/model"
	"github.com/nkarpenko/koho-transaction/conf"
	"github.com/nkarpenko/koho-transaction/validator"
)

// Transaction interface holds a collection of methods that help validate user
// transactions.
type Transaction interface {
	Process(*model.Result) error
	Validate(*model.Transaction) *model.Result
}

// transaction struct holds a collection of required interfaces for the
// transaction service.
type transaction struct {
	validator validator.Validator
}

// Process function stores the transaction data to later check
// against.
func (t *transaction) Process(res *model.Result) error {

	// Add transaction to cache.
	(*cache.Cache)[res.CustomerID] = append((*cache.Cache)[res.CustomerID], *res)

	// Sort cache key values by date.
	sort.Slice((*cache.Cache)[res.CustomerID], func(i, j int) bool {
		return (*cache.Cache)[res.CustomerID][i].Time.After((*cache.Cache)[res.CustomerID][j].Time)
	})

	// Convert the result to a json string.
	json, err := json.Marshal(res)
	if err != nil {
		return err
	}

	// Output the final results. At this point we can use the transaction struct
	// settings to point it to the db, redis, or return it to the user via api.
	fmt.Println(string(json))

	// Successful validation.
	return nil
}

// Validate method validates a users transaction to make sure they are within
// their transaction limits.
func (t *transaction) Validate(tx *model.Transaction) *model.Result {

	// Validate the transaction via validator interface.
	return t.validator.Validate(tx)
}

// New transaction service instance.
func New(c *conf.Config) Transaction {
	return &transaction{
		validator: validator.New(c),
	}
}
