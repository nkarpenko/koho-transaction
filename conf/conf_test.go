package conf

import (
	"testing"
)

type test struct {
	result bool
	file   string
}

func TestLoad(t *testing.T) {

	// Initialize test cases.
	tests := []test{
		{
			result: false,
			file:   "./input.txt",
		},
		{
			result: false,
			file:   "./invalid_input.txt",
		},
		{
			result: true,
			file:   "../input.txt",
		},
	}

	// Run test cases.
	for _, test := range tests {
		conf, err := Load(test.file)
		if test.result && err != nil {
			t.Errorf("unable to initialize app: %+v", err)
		}

		if test.result && conf == nil {
			t.Errorf("unable to load taml configuration file: %+v", err)
		}
	}
}
