package transaction

import (
	"testing"

	"github.com/nkarpenko/koho-transaction/common/model"
	"github.com/nkarpenko/koho-transaction/conf"
)

type test struct {
	result       bool
	config       *conf.Config
	transactions []model.Transaction
	results      []model.Result
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

func TestProcess(t *testing.T) {

}

func TestValidate(t *testing.T) {

}
