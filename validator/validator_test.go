package validator

import (
	"testing"
	"time"

	"github.com/nkarpenko/koho-transaction/common/model"
	"github.com/nkarpenko/koho-transaction/conf"
)

type test struct {
	result       bool
	results      []bool
	config       *conf.Config
	txid         int
	accepted     bool
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

	// Run test cases.
	for _, test := range tests {
		v := New(test.config)
		if v == nil && test.result {
			t.Error("validator not initialized")
		}
	}
}

func TestValidate(t *testing.T) {

	// Initialize test cases.
	tests := []test{
		{
			results: []bool{
				true,
				true,
				false,
			},
			transactions: []model.Transaction{
				{
					ID:         1,
					CustomerID: 1,
					LoadAmount: 1750,
					Time:       time.Now(),
				},
				{
					ID:         2,
					CustomerID: 2,
					LoadAmount: 3600,
					Time:       time.Now(),
				},
				{
					ID:         3,
					CustomerID: 3,
					LoadAmount: 5001,
					Time:       time.Now(),
				},
			},
			config: &conf.Config{
				Name: "Test valid config",
				Limits: &model.Limits{
					DailyAmount:       5000,
					DailyTransactions: 3,
					WeeklyAmount:      20000,
				},
			},
		},
	}

	// Run test cases.
	for i, test := range tests {
		v := New(test.config)

		res := v.Validate(&test.transactions[i])
		if res.Accepted != test.results[i] {
			t.Errorf("test case id '%d' excpected value of '%+v' does not match expected value %+v'", test.transactions[i].ID, test.results[i], res.Accepted)
		}
		if v == nil && test.result {
		}
	}
}

func TestIsUniqueTransactionID(t *testing.T) {

	// Initialize test cases.
	tests := []test{
		{
			result: false,
			txid:   7,
			transactions: []model.Transaction{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
				{
					ID: 1,
				},
			},
			config: &conf.Config{
				Limits: &model.Limits{
					DailyAmount:       5000,
					DailyTransactions: 3,
					WeeklyAmount:      20000,
				},
			},
		},
		{
			result: true,
			txid:   1,
			transactions: []model.Transaction{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
				{
					ID: 3,
				},
			},
			config: &conf.Config{
				Limits: &model.Limits{
					DailyAmount:       5000,
					DailyTransactions: 3,
					WeeklyAmount:      20000,
				},
			},
		},
	}

	// Run test cases.
	for _, test := range tests {
		v := New(test.config)

		for _, tx := range test.transactions {

			res := v.Validate(&tx)
			if test.txid == res.ID && !test.result {
				t.Error("expected test case to pass")
			}
		}
	}
}

func TestIsWithinDailyAmountLimit(t *testing.T) {

	// Initialize test cases.
	tests := []test{
		{
			accepted: true,
			transactions: []model.Transaction{
				{
					ID:         1,
					CustomerID: 1,
					LoadAmount: 2600,
					Time:       time.Now(),
				},
				{
					ID:         2,
					CustomerID: 1,
					LoadAmount: 2600,
					Time:       time.Now(),
				},
			},
			config: &conf.Config{
				Name: "Test valid config",
				Limits: &model.Limits{
					DailyAmount:       6000,
					DailyTransactions: 3,
					WeeklyAmount:      20000,
				},
			},
		},
		{
			accepted: true,
			transactions: []model.Transaction{
				{
					ID:         1,
					CustomerID: 1,
					LoadAmount: 2500,
					Time:       time.Now(),
				},
				{
					ID:         2,
					CustomerID: 1,
					LoadAmount: 2499,
					Time:       time.Now(),
				},
			},
			config: &conf.Config{
				Name: "Test valid config",
				Limits: &model.Limits{
					DailyAmount:       5000,
					DailyTransactions: 3,
					WeeklyAmount:      20000,
				},
			},
		},
	}

	// Run test cases.
	for _, test := range tests {
		v := New(test.config)

		for _, tx := range test.transactions {

			res := v.Validate(&tx)
			if test.accepted != res.Accepted {
				t.Error("expected test case to pass")
			}
		}
	}
}

func TestIsWithinDailyLoadLimit(t *testing.T) {

	// Initialize test cases.
	tests := []test{
		{
			accepted: false,
			transactions: []model.Transaction{
				{
					ID:         1,
					CustomerID: 1,
					LoadAmount: 10000,
					Time:       time.Now(),
				},
				{
					ID:         2,
					CustomerID: 1,
					LoadAmount: 15000,
					Time:       time.Now(),
				},
			},
			config: &conf.Config{
				Name: "Test valid config",
				Limits: &model.Limits{
					DailyAmount:       5000,
					DailyTransactions: 3,
					WeeklyAmount:      20000,
				},
			},
		},
		{
			accepted: true,
			transactions: []model.Transaction{
				{
					ID:         1,
					CustomerID: 1,
					LoadAmount: 2500,
					Time:       time.Now(),
				},
				{
					ID:         2,
					CustomerID: 1,
					LoadAmount: 2499,
					Time:       time.Now(),
				},
			},
			config: &conf.Config{
				Name: "Test valid config",
				Limits: &model.Limits{
					DailyAmount:       5000,
					DailyTransactions: 3,
					WeeklyAmount:      20000,
				},
			},
		},
	}

	// Run test cases.
	for _, test := range tests {
		v := New(test.config)

		for _, tx := range test.transactions {

			res := v.Validate(&tx)
			if test.accepted != res.Accepted {
				t.Error("expected test case to pass")
			}
		}
	}
}

func TestIsWithinWeeklyAmountLimit(t *testing.T) {

	// Initialize test cases.
	tests := []test{
		{
			accepted: false,
			transactions: []model.Transaction{
				{
					ID:         1,
					CustomerID: 1,
					LoadAmount: 10000,
					Time:       time.Now(),
				},
				{
					ID:         2,
					CustomerID: 1,
					LoadAmount: 15000,
					Time:       time.Now(),
				},
			},
			config: &conf.Config{
				Name: "Test valid config",
				Limits: &model.Limits{
					DailyAmount:       5000,
					DailyTransactions: 3,
					WeeklyAmount:      20000,
				},
			},
		},
		{
			accepted: true,
			transactions: []model.Transaction{
				{
					ID:         1,
					CustomerID: 1,
					LoadAmount: 2500,
					Time:       time.Now(),
				},
				{
					ID:         2,
					CustomerID: 1,
					LoadAmount: 2499,
					Time:       time.Now(),
				},
			},
			config: &conf.Config{
				Name: "Test valid config",
				Limits: &model.Limits{
					DailyAmount:       5000,
					DailyTransactions: 3,
					WeeklyAmount:      20000,
				},
			},
		},
	}

	// Run test cases.
	for _, test := range tests {
		v := New(test.config)

		for _, tx := range test.transactions {

			res := v.Validate(&tx)
			if test.accepted != res.Accepted {
				t.Error("expected test case to pass")
			}
		}
	}
}
