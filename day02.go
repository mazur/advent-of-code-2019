package main

import "fmt"

func IntcodeComputer(p []int) []int {
  for i := 0; i < len(p); i=i+4 {
    if p[i] == 99 {
      return p
    }

    if p[i] == 1 {
      p[p[i+3]] = p[p[i+1]] + p[p[i+2]]
    } else if p[i] == 2 {
      p[p[i+3]] = p[p[i+1]] * p[p[i+2]]
    } else {
      break
    }

  }
  return p
}

func GetNewMemory() []int {
    return []int {1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,9,19,1,5,19,23,1,6,23,27,1,27,10,31,1,31,5,35,2,10,35,39,1,9,39,43,1,43,5,47,1,47,6,51,2,51,6,55,1,13,55,59,2,6,59,63,1,63,5,67,2,10,67,71,1,9,71,75,1,75,13,79,1,10,79,83,2,83,13,87,1,87,6,91,1,5,91,95,2,95,9,99,1,5,99,103,1,103,6,107,2,107,13,111,1,111,10,115,2,10,115,119,1,9,119,123,1,123,9,127,1,13,127,131,2,10,131,135,1,135,5,139,1,2,139,143,1,143,5,0,99,2,0,14,0}
}

func FindNounAndVerb() (int, int) {
  var noun, verb int

  for noun = 0; noun < 100; noun++ {
    for verb = 0; verb < 100; verb++ {
      p := GetNewMemory()
      p[1] = noun
      p[2] = verb
      output := IntcodeComputer(p)
      if output[0] == 19690720 {
        return noun, verb
      }
    }
  }
  return -1, -1
}

func Day02() string {
  noun, verb := FindNounAndVerb()
  awnser := 100 * noun + verb

  return fmt.Sprintf("%d", awnser)
}
