package main

import "testing"

func TestFuelCalculator(t *testing.T) {
  testCases := []struct {
    mass int
		fuel int
	}{
		{1969, 966},
		{100756, 50346},
	}

	for _, testCase := range testCases {
		fuel := FuelCalculator(testCase.mass)
		if fuel != testCase.fuel {
			t.Errorf("Fuel for mass %d was incorrect, got: %d, want: %d.", testCase.mass, fuel, testCase.fuel)
		}
	}
}
