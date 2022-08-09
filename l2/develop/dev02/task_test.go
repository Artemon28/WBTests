package main

import (
	"errors"
	"strings"
	"testing"
)

func TestStringUnpacking(t *testing.T) {

	testTable := []struct {
		initString     string
		expectedString string
		err            error
	}{
		{
			initString:     "a4bc2d5e",
			expectedString: "aaaabccddddde",
			err:            nil,
		},
		{
			initString:     "abcd",
			expectedString: "abcd",
			err:            nil,
		},
		{
			initString:     "",
			expectedString: "",
			err:            nil,
		},
		{
			initString:     "45",
			expectedString: "",
			err:            errors.New(""),
		},
		{
			initString:     "4sdg5sgesf4t4s4g4",
			expectedString: "",
			err:            errors.New(""),
		},
		{
			initString:     "q1",
			expectedString: "q",
			err:            nil,
		},
		{
			initString:     "q6",
			expectedString: "qqqqqq",
			err:            nil,
		},
	}

	for _, testCase := range testTable {
		resultString, err := StringUnpacking(testCase.initString)

		if err != nil && testCase.err == nil {
			t.Error(err.Error())
		}

		if strings.Compare(resultString, testCase.expectedString) != 0 {
			t.Errorf("this strings should be the same, but they different\nresult string: %s\nexpected string: %s",
				resultString, testCase.expectedString)
		}
	}

}
