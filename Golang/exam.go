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
	var fl float64
	var num int64

	fmt.Print("Enter a floating point number : ")
	fmt.Scanf("%f", &fl)

	num = int64(fl)
	fmt.Printf(" Float %f Integer %d \n", fl, num);
}
