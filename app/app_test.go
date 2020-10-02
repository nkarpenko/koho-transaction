package app

import (
	"testing"

	"github.com/nkarpenko/koho-transaction/common/model"
	"github.com/nkarpenko/koho-transaction/conf"
)

type test struct {
	result bool
	config *conf.Config
}

func TestStart(t *testing.T) {

	// Initialize test cases.
	tests := []test{
		{
			config: &conf.Config{
				Name:      "Test Conf 1",
				InputFile: "./input.txt",
				Limits: &model.Limits{
					DailyAmount:       6000,
					DailyTransactions: 5,
					WeeklyAmount:      30000,
				},
			},
		},
		{
			config: &conf.Config{
				Name:      "Test Conf 2",
				InputFile: "./input.txt",
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
		app, err := New(test.config)
		if err != nil {
			t.Errorf("unable to initialize app: %+v", err)
		}

		app.Start()
	}

}

func TestNew(t *testing.T) {

	// Initialize test cases.
	tests := []test{
		{
			config: &conf.Config{
				Name:      "Test Conf 1",
				InputFile: "./input.txt",
				Limits: &model.Limits{
					DailyAmount:       6000,
					DailyTransactions: 5,
					WeeklyAmount:      30000,
				},
			},
		},
		{
			config: &conf.Config{
				Name:      "Test Conf 2",
				InputFile: "./input.txt",
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
		_, err := New(test.config)
		if err != nil {
			t.Errorf("unable to initialize app: %+v", err)
		}
	}
}
