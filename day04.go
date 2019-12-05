package main

import (
    "strconv"
    "fmt"
)

func ValidatePassword(password int) bool {
  passstr := strconv.Itoa(password)
  prev := -1
  count := 0
  double := false

  for i:=0; i < len(passstr); i++ {
    curr, _ := strconv.Atoi(string(passstr[i]))
    if curr < prev {
      return false
    } else if prev == curr {
      count++;
    } else {
      double = double || (count == 1)
      count = 0
    }

    prev = curr
  }

  return double || (count == 1)
}

func NumberOfCorrectPasswords(start int, end int) int {
  count := 0
  for i := start; i < end; i++ {
    if ValidatePassword(i) {
      count++
    }
  }
  return count
}

func Day04() string {
  awnser := NumberOfCorrectPasswords(265275,781584)

  return fmt.Sprintf("%d", awnser)
}
