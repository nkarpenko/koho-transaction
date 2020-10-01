package model

import (
	"encoding/json"
	"strconv"
	"time"
)

// Transaction struct contains a user's single transaction request.
type Transaction struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	LoadAmount float64   `json:"load_amount"`
	Time       time.Time `json:"time"`
}

// Result struct holds the validation results and transaction data to process.
type Result struct {
	ID         int    `json:"id"`
	CustomerID int    `json:"customer_id"`
	Accepted   bool   `json:"accepted"`
	Reason     string `json:"-"` // enable json field for debugging

	// Don't print these but keep them for cache purposes.
	LoadAmount float64   `json:"-"`
	Time       time.Time `json:"-"`
}

// Output struct contains the vars and converted types for the final application output.
type Output struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Accepted   bool   `json:"accepted"`
	Reason     string `json:"-"` // enable json field for debugging
}

// Limits struct holds details on user transaction limits.
type Limits struct {
	DailyAmount       int `mapstructure:"daily_amount"`
	DailyTransactions int `mapstructure:"daily_transactions"`
	WeeklyAmount      int `mapstructure:"weekly_amount"`
}

// UnmarshalJSON implements a custom scanner for the transaction type.
func (t *Transaction) UnmarshalJSON(b []byte) error {

	var (
		err error
		v   map[string]interface{}
	)

	// Unmarshal transaction into map first before manual type conversion.
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	// Convert id string to int.
	if id, ok := v["id"].(string); ok {

		// Set transaction ID to integer value.
		t.ID, err = strconv.Atoi(id)
		if err != nil {
			return err
		}
	}

	// Convert customer id string to int.
	if cid, ok := v["customer_id"].(string); ok {

		// Set customer ID to integer value.
		t.CustomerID, err = strconv.Atoi(cid)
		if err != nil {
			return err
		}
	}

	// Convert customer id string to float64.
	if amt, ok := v["load_amount"].(string); ok {

		// If first character is a dollar sign $ remove it before type conversion.
		if amt[0:1] == "$" {
			amt = amt[1:]
		}

		// Set load amount to float value.
		t.LoadAmount, err = strconv.ParseFloat(amt, 64)
		if err != nil {
			return err
		}
	}

	// Convert time string to time.Time.
	if date, ok := v["time"].(string); ok {

		// Convert time to time.Time value.
		format := "2006-01-02T15:04:05Z"
		t.Time, err = time.Parse(format, date)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON implements a custom scanner for the final result object
func (r *Result) MarshalJSON() ([]byte, error) {

	// if the result struct is not valid, set it as null in the json
	if r == nil {
		return []byte("{}"), nil
	}

	// Set the converted values.
	var res = &Output{}
	res.ID = strconv.Itoa(r.ID)
	res.CustomerID = strconv.Itoa(r.CustomerID)
	res.Accepted = r.Accepted
	res.Reason = r.Reason // for debugging enable it in the structs.

	// Final conversion to json string.
	b, err := json.Marshal(res)
	return b, err
}
