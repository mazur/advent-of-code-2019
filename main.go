package main

import (
	"fmt"
	"os"
	"strconv"
)

type DayFunc func() string

func main() {
	days := []DayFunc{
		Day01,
		Day02,
	}

	if len(os.Args) == 1 {
		for i := 0; i < len(days); i++ {
			ExecDay(days[i], i+1)
		}
		return
	}

	dayString := os.Args[1]
	i, err := strconv.Atoi(dayString)

	if err != nil {
		fmt.Printf("Could not parse %q to day number.\n", dayString)
		return;
	}
	if i <= 0 {
		fmt.Printf("Day can't be negative or zero.\n")
		return
	}
	if i > len(days) {
		fmt.Printf("Day%02d can't be found.\n", i)
		return
	}

	ExecDay(days[i-1], i)
}

func ExecDay(dayFunc DayFunc, day int) {
	fmt.Printf("Day%02d: %s\n", day, dayFunc())
}
