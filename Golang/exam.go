package main

import (
	"fmt"
)

/*
Write a program which prompts the user to enter a floating point number and prints the integer
which is a truncated version of the floating point number that was entered.
Truncation is the process of removing the digits to the right of the decimal place.
*/
func main() {
  s := make([]int, 0, 3)
  s = append(s, 100)
  fmt.Println(len(s), cap(s))
}
