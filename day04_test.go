package main

import "testing"

func TestValidatePassword(t *testing.T) {
  testCases := []struct {
    password int
		valid bool
	}{
    {112233, true},
    {111111, false},
    {223450, false},
    {123789, false},
    {112233, true},
    {123444, false},
    {111122, true},
    {112344, true},
    {113444, true},
	}

	for _, testCase := range testCases {
		output := ValidatePassword(testCase.password)
		if output != testCase.valid {
			t.Errorf("Password validity for password %d was incorrect, got: %v, want: %v.", testCase.password, output, testCase.valid)
		}
	}
}
