package transaction

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/nkarpenko/koho-transaction/conf"
	"github.com/nkarpenko/koho-transaction/transaction/model"
)

// Cache is the cache for storing transaction data. In a real production scenario,
// we would use some memory store caching mechanism (e.g. Redis) or db.
var Cache = &map[int][]model.Result{}

// Validator interface holds a collection of methods
// that help validate user transactions.
type Validator interface {
	// Validation methods
	Validate(*model.Transaction) *model.Result

	// Transction methods.
	ProcessTransaction(res *model.Result) error

	// Bool methods.
	IsUniqueTransactionID(customerID int, txid int) bool
	IsWithinDailyLoadLimit(customerID int, date time.Time) bool
	IsWithinDailyAmountLimit(customerID int, date time.Time, amount float64) bool
	IsWithinWeeklyAmountLimit(customerID int, date time.Time, amount float64) bool
}

// validator struct holds a collection of required vars
// and config for the Validator interface.
type validator struct {
	limits *model.Limits
}

// Validate method validates a users transaction to make
// sure they are within their transaction limits.
func (v *validator) Validate(tx *model.Transaction) *model.Result {

	// Init new result model.
	var res = &model.Result{}

	// Set default values.
	res.ID = tx.ID
	res.CustomerID = tx.CustomerID
	res.LoadAmount = tx.LoadAmount
	res.Time = tx.Time
	res.Accepted = true

	// Confirm the transaction id is unique and hasn't been already processed.
	// Return early if a duplicate id exists and it is not unique.
	if res.Accepted = v.IsUniqueTransactionID(res.CustomerID, res.ID); !res.Accepted {
		res.Reason = "id not unique for customer"
		return res
	}

	// Confirm user is within load limits and return false if they are not.
	if res.Accepted = v.IsWithinDailyLoadLimit(res.CustomerID, res.Time); !res.Accepted {
		res.Reason = "daily load limit exceeded"
		return res
	}

	// Confirm user is within daily load amount limit and return early if not.
	if res.Accepted = v.IsWithinDailyAmountLimit(res.CustomerID, res.Time, res.LoadAmount); !res.Accepted {
		res.Reason = "daily amount limit exceeded"
		return res
	}

	// Confirm user is within weekly load amount limit and return early if not.
	if res.Accepted = v.IsWithinWeeklyAmountLimit(res.CustomerID, res.Time, res.LoadAmount); !res.Accepted {
		res.Reason = "weekly amount limit exceeded"
		return res
	}

	return res
}

// ProcessTransaction function stores the transaction data to later check against.
func (v *validator) ProcessTransaction(res *model.Result) error {

	// Add transaction to cache.
	(*Cache)[res.CustomerID] = append((*Cache)[res.CustomerID], *res)

	// Sort cache values by date.
	sort.Slice((*Cache)[res.CustomerID], func(i, j int) bool { return (*Cache)[res.CustomerID][i].Time.After((*Cache)[res.CustomerID][j].Time) })

	// Convert the result to a json string.
	json, err := json.Marshal(res)
	if err != nil {
		return err
	}

	// Output the final results.
	fmt.Println(string(json))

	return nil
}

func (v *validator) IsUniqueTransactionID(cid int, txid int) (accepted bool) {

	// Check if cache data already exists for this customer.
	// Return true if it doesn't meaning this is their first load
	// so transaction id MUST be unique for them.
	data, ok := (*Cache)[cid]
	if !ok {
		return true
	}

	// Loop through cache to check if ID is unique.
	for _, entry := range data {

		// If the new transaction ID matches an ID in the cache, don't accept it.
		if entry.ID == txid {
			return false
		}
	}

	// Successfully validated and accepted.
	return true
}

func (v *validator) IsWithinDailyLoadLimit(customerID int, date time.Time) (accepted bool) {

	// Check if cache data already exists for this customer.
	// Return true if it doesn't meaning this is their first load.
	data, ok := (*Cache)[customerID]
	if !ok {
		return true
	}

	// Init counter and 24 hour date limit.
	count := 0
	limit := timeToDayStart(date)

	// Loop through cache entries to count amount of transaction for the user.
	for _, entry := range data {

		// Add count if an entry is between 24 hour agos and now.
		if entry.Time.After(limit) && entry.Time.Before(date) {
			count++
		}
	}

	// Compare total count of loads to validator limit
	if count >= v.limits.DailyTransactions {
		return false
	}

	// Successfully validated and accepted.
	return true
}

func (v *validator) IsWithinDailyAmountLimit(customerID int, date time.Time, amount float64) (accepted bool) {

	// Get the cached user transaction data.
	data, _ := (*Cache)[customerID]

	// Init load amount counter and 24 hour date limit.
	limit := timeToDayStart(date)

	// Loop through cache entries to count amount of transaction for the user.
	for _, entry := range data {

		// Add if entry was accepted and count if an entry is between 24 hour agos and now.
		if entry.Accepted && entry.Time.After(limit) && entry.Time.Before(date) {

			// Increment the total amount.
			amount = amount + entry.LoadAmount
		}
	}

	// Compare total load amount for last day to validator limit
	if amount >= float64(v.limits.DailyAmount) {

		// User has exceeded allowed daily amount.
		return false
	}

	// Successfully validated and accepted.
	return true
}

func (v *validator) IsWithinWeeklyAmountLimit(customerID int, date time.Time, amount float64) (accepted bool) {

	// Get the cached user transaction data.
	data, _ := (*Cache)[customerID]

	// Init load amount counter and 1 week date limit (starting monday).
	limit := timeToWeekStart(date)

	// Loop through cache entries to count amount of transaction for the user.
	for _, entry := range data {

		// Add count entry was accepted and if an entry is between a wee ago and now.
		if entry.Accepted && entry.Time.After(limit) && entry.Time.Before(date) {
			amount = amount + entry.LoadAmount
		}
	}

	// Compare total load amount for last day to validator limit
	if amount >= float64(v.limits.WeeklyAmount) {
		return false
	}

	// Successfully validated and accepted.
	return true
}

func timeToDayStart(t time.Time) time.Time {
	year, month, day := t.Date()

	// End day 1 second before midnight
	return time.Date(year, month, day, 0, 0, -1, 0, t.Location())
}

func timeToWeekStart(t time.Time) time.Time {
	year, month, day := t.Date()
	date := time.Date(year, month, day, 0, 0, 0, 0, t.Location())

	// Keep iterating back until date is the start of monday.
	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, -1)
	}

	return date
}

// NewValidator instance.
func NewValidator(c *conf.Config) Validator {
	return &validator{
		limits: c.Limits,
	}
}
