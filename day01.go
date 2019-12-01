package main

import "fmt"

func FuelCalculator(mass int) int {
  fuel := (mass / 3) - 2

  if fuel <= 0 {
    return 0
  }

  return fuel + FuelCalculator(fuel)
}

func CalculateSumOfFuel(masses []int) int {
  sum := 0
  for i := 0; i < len(masses); i++ {
		sum += FuelCalculator(masses[i])
	}
  return sum
}

func Day01() string {
  inputs := []int{
    64010,
    104993,
    95523,
    87818,
    88717,
    107983,
    145542,
    105501,
    114620,
    58641,
    103006,
    143491,
    50414,
    112904,
    87089,
    128212,
    65482,
    55270,
    135648,
    104915,
    82940,
    117582,
    140160,
    108526,
    89334,
    56984,
    56359,
    92300,
    93287,
    122020,
    121921,
    58083,
    78671,
    115880,
    63348,
    59915,
    124435,
    93727,
    100850,
    121204,
    58303,
    70936,
    106085,
    101849,
    88080,
    136756,
    146614,
    126070,
    108147,
    55586,
    72262,
    74494,
    113382,
    139642,
    91326,
    109113,
    104840,
    112594,
    123700,
    130201,
    135021,
    75904,
    148338,
    117256,
    92802,
    86456,
    124484,
    127723,
    53713,
    55862,
    120367,
    77904,
    138061,
    65166,
    135541,
    109393,
    102805,
    131760,
    130030,
    114953,
    101461,
    72993,
    146507,
    112431,
    145712,
    139519,
    122758,
    80609,
    91775,
    73807,
    77878,
    112319,
    110665,
    119908,
    124568,
    67409,
    123830,
    130549,
    114312,
    79899,
  }

  return fmt.Sprintf("%d", CalculateSumOfFuel(inputs))
}
