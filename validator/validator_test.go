package validator

import (
	"testing"

	"github.com/nkarpenko/koho-transaction/common/model"
	"github.com/nkarpenko/koho-transaction/conf"
)

type test struct {
	result       bool
	config       *conf.Config
	transactions []model.Transaction
}

func TestNew(t *testing.T) {

	// Initialize test cases.
	tests := []test{
		{
			result: true,
			config: &conf.Config{
				Name: "Test Conf 1",
				Limits: &model.Limits{
					DailyAmount:       6000,
					DailyTransactions: 5,
					WeeklyAmount:      30000,
				},
			},
		},
		{
			result: true,
			config: &conf.Config{
				Name: "Test Conf 2",
				Limits: &model.Limits{
					DailyAmount:       1000,
					DailyTransactions: 2,
					WeeklyAmount:      10000,
				},
			},
		},
	}

	for _, test := range tests {
		v := New(test.config)
		if v == nil && test.result {
			t.Error("validator not initialized")
		}
	}
}

func TestValidate(t *testing.T) {

}

func TestIsUniqueTransactionID(t *testing.T) {

}

func TestIsWithinDailyAmountLimit(t *testing.T) {

}

func TestIsWithinDailyLoadLimit(t *testing.T) {

}

func TestIsWithinWeeklyAmountLimit(t *testing.T) {

}

// Validate(*model.Transaction) *model.Result
// IsUniqueTransactionID(customerID int, txid int) bool
// IsWithinDailyAmountLimit(customerID int, date time.Time, amount float64) bool
// IsWithinDailyLoadLimit(customerID int, date time.Time) bool
// IsWithinWeeklyAmountLimit(customerID int, date time.Time, amount float64) bool
// New
