package main

import "testing"

func TestDistanceToCentralPort(t *testing.T) {
  testCases := []struct {
    wires []string
		distance int
	}{
    {[]string{"U1,R1","R2,D1,L1,U2"}, 2},
    {[]string{"R8,U5,L5,D3", "U7,R6,D4,L4"}, 6},
    {[]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}, 159},
		{[]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51","U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}, 135},
	}

	for _, testCase := range testCases {
		distance := DistanceToCentralPort(testCase.wires)
		if distance != testCase.distance {
			t.Errorf("%v\n\nDistance for wires was incorrect, got: %d, want: %d.", testCase.wires, distance, testCase.distance)
		}
	}
}
