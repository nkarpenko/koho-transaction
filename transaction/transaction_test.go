package transaction

import (
	"testing"

	"github.com/nkarpenko/koho-transaction/common/model"
	"github.com/nkarpenko/koho-transaction/conf"
)

type test struct {
	result  bool
	config  *conf.Config
	results model.Result
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
		tx := New(test.config)
		if tx == nil && test.result {
			t.Error("transaction service not initialized")
		}
	}
}

func TestProcess(t *testing.T) {
	// Initialize test cases.
	tests := []test{
		{
			result: true,
			results: model.Result{
				ID:         1,
				CustomerID: 2,
				Accepted:   true,
			},
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
			results: model.Result{
				ID:         2,
				CustomerID: 3,
				Accepted:   true,
			},
			config: &conf.Config{
				Name: "Test Conf 2",
				Limits: &model.Limits{
					DailyAmount:       1000,
					DailyTransactions: 2,
					WeeklyAmount:      10000,
				},
			},
		},
		{
			result: true,
			results: model.Result{
				ID:         4,
				CustomerID: 5,
				Accepted:   true,
			},
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
		tx := New(test.config)
		err := tx.Process(&test.results)
		if err != nil {
			t.Errorf("unable to process transaction: %+v", err)
		}
	}
}
