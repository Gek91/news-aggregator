package main

import (
	"slices"
	"testing"
)

func testRetrieveFilterFromString(t *testing.T) {

	var tests = []struct {
		input  string
		result []string
	}{
		{"", nil},
		{" ", nil},
		{"value", []string{"value"}},
		{"value1,value2", []string{"value1", "value2"}},
	}

	for _, test := range tests {

		if result := retrieveFilterFromString(test.input); !slices.Equal(result, test.result) {
			t.Errorf("retrieveFilterFromString(%q) result not %v", test.input, test.result)
		}

	}

}
