package parser

import (
	"testing"

	"github.com/nkarpenko/koho-transaction/common/model"
	"github.com/nkarpenko/koho-transaction/conf"
)

type test struct {
	result        bool
	config        *conf.Config
	inputFilePath string
}

func TestParseFile(t *testing.T) {

	// Initialize test cases.
	tests := []test{
		{
			result: false,
			config: &conf.Config{
				Name:      "Test Conf 1",
				InputFile: "./invalid_path.txt",
				Limits: &model.Limits{
					DailyAmount:       5000,
					DailyTransactions: 3,
					WeeklyAmount:      20000,
				},
			},
		},
		{
			result: true,
			config: &conf.Config{
				Name:      "Test Conf 2",
				InputFile: "../input.txt",
				Limits: &model.Limits{
					DailyAmount:       6000,
					DailyTransactions: 5,
					WeeklyAmount:      30000,
				},
			},
		},
	}

	// Run test cases.
	for _, test := range tests {
		p := New(test.config)
		_, err := p.ParseFile()
		if err != nil && test.result {
			t.Errorf("failed to parse file: %+v", err)
		}
	}
}

func TestNew(t *testing.T) {

	// Initialize test cases.
	tests := []test{
		{
			result: true,
			config: &conf.Config{
				Name:      "Test Conf 1",
				InputFile: "../input.txt",
			},
		},
		{
			result: true,
			config: &conf.Config{
				Name:      "Test Conf 2",
				InputFile: "../input.txt",
			},
		},
	}

	// Run test cases.
	for _, test := range tests {
		p := New(test.config)
		if p == nil && test.result {
			t.Error("validator not initialized")
		}
	}
}
