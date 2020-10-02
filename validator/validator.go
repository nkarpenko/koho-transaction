// Package validator contains a collection of interfaces and methods required
//  to validation user transaction requests.
package validator

import (
	"time"

	"github.com/nkarpenko/koho-transaction/common/cache"
	"github.com/nkarpenko/koho-transaction/common/model"
	"github.com/nkarpenko/koho-transaction/conf"
)

// Validator interface holds a collection of methods to validate any incoming
// user transaction requests.
type Validator interface {
	// Main validation methods.
	Validate(*model.Transaction) *model.Result

	// Bool methods.
	IsUniqueTransactionID(customerID int, txid int) bool
	IsWithinDailyAmountLimit(customerID int, date time.Time, amount float64) bool
	IsWithinDailyLoadLimit(customerID int, date time.Time) bool
	IsWithinWeeklyAmountLimit(customerID int, date time.Time, amount float64) bool
}

// validator struct holds a collection of config vars required for various
// validation methods.
type validator struct {
	limits *model.Limits
}

// Validate method validates a users transaction to make sure they are within
// their transaction limits supplied in the configuration.
func (v *validator) Validate(tx *model.Transaction) *model.Result {

	// Init new result transaction model.
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
		// Provide failure Message for debug.
		res.Message = "transaction id is not unique for customer, ignoring"
		return res
	}

	// Confirm user is within load limits and return false if they are not.
	if res.Accepted = v.IsWithinDailyLoadLimit(res.CustomerID, res.Time); !res.Accepted {
		// Provide failure Message for debug.
		res.Message = "daily load limit exceeded"
		return res
	}

	// Confirm user is within daily load amount limit and return early if not.
	if res.Accepted = v.IsWithinDailyAmountLimit(res.CustomerID, res.Time, res.LoadAmount); !res.Accepted {
		// Provide failure Message for debug.
		res.Message = "daily amount limit exceeded"
		return res
	}

	// Confirm user is within weekly load amount limit and return early if not.
	if res.Accepted = v.IsWithinWeeklyAmountLimit(res.CustomerID, res.Time, res.LoadAmount); !res.Accepted {
		// Provide failure Message for debug.
		res.Message = "weekly amount limit exceeded"
		return res
	}

	return res
}

// IsUniqueTransactionID method validates the transtion ID is unique to
// the specified user.
func (v *validator) IsUniqueTransactionID(cid int, txid int) (accepted bool) {

	// Check if cache data already exists for this customer. Return true if the
	// key doesn't exist since it means it's their first transaction.
	data, ok := (*cache.Cache)[cid]
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

// IsWithinDailyAmountLimit validates that the user's daily load amount is
// within its daily limit specified inside of the config.yml file.
func (v *validator) IsWithinDailyAmountLimit(customerID int, date time.Time, amount float64) (accepted bool) {

	// Get the cached user transaction data. Ignore if it DNE since we will have
	//  to add new data to it anyway.
	data, _ := (*cache.Cache)[customerID]

	// Init load amount counter and 24 hour date limit.
	limit := timeToDayStart(date)

	// Loop through cache entries to count amount of transaction for the user.
	for _, entry := range data {

		// Add if entry was accepted and count if an entry is between 24 hour agos
		// and now.
		if entry.Accepted && entry.Time.After(limit) && entry.Time.Before(date) {

			// Increment the total amount.
			amount = amount + entry.LoadAmount
		}
	}

	// Compare total load amount for last day to validator limit.
	if amount >= float64(v.limits.DailyAmount) {

		// User has exceeded allowed daily amount.
		return false
	}

	// Successfully validated and accepted.
	return true
}

// IsWithinDailyLoadLimit validates that the user's daily transaction count is
// within its daily limit specified inside of the config.yml file.
func (v *validator) IsWithinDailyLoadLimit(customerID int, date time.Time) (accepted bool) {

	// Check if cache data already exists for this customer.
	// Return true if it doesn't meaning this is their first transaction.
	data, ok := (*cache.Cache)[customerID]
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

// IsWithinWeeklyAmountLimit validates that the user's weekly load amount is
// within its daily limit specified inside of the config.yml file.
func (v *validator) IsWithinWeeklyAmountLimit(customerID int, date time.Time, amount float64) (accepted bool) {

	// Get the cached user transaction data.
	data, _ := (*cache.Cache)[customerID]

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

// timeToDayStart helper method returns the start date/time of the specific day
// given (Midnight -1 second).
func timeToDayStart(t time.Time) time.Time {
	year, month, day := t.Date()

	// End day 1 second before midnight
	return time.Date(year, month, day, 0, 0, -1, 0, t.Location())
}

// timeToWeekStart helper method returns the date for the start of the week
// (Monday Midnight)
func timeToWeekStart(t time.Time) time.Time {
	year, month, day := t.Date()
	date := time.Date(year, month, day, 0, 0, 0, 0, t.Location())

	// Keep iterating back until date is the start of monday.
	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, -1)
	}

	return date
}

// New Validator instance.
func New(c *conf.Config) Validator {
	return &validator{
		limits: c.Limits,
	}
}
