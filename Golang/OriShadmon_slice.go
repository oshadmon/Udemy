/**
Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should
be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The
program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice
must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit
(exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

package main
import (
    "fmt"
    "math"
    "sort"
    "strconv"
)

func __convert_string(user_input string) int {
    /**
        Given a user value convert into int
        :args:
            user_input:string - user input as a string
        :param:
            number:int - user value as an integer
        :return:
            nunber
    */
    number,_ := strconv.ParseUint(user_input, 10, 32)
    return int(number)

}


func main() {
    /**
        User inputs value into list, then list is sorted and printed
        :params:
            slice:list - growing array
            user_input:str - user input as string
            int_user_input:int - converted user input to integer
            i:int - placement on the growing array
    */
    slice := make([]int, 0, math.MaxInt32) // slice with indefinite max size
    var user_input string
    var i int = 0

    for true {
        fmt.Printf("Input Number: ")
        fmt.Scan(&user_input)
        if user_input == "x" || user_input == "X" {
            // check whether user inputted x - if so exit
            break
        }

        // convert string to integer
        int_user_input := __convert_string(user_input)

        // insert value
        slice = append(slice, int_user_input)

        // sort slice
        sort.Sort(sort.IntSlice(slice))

        // print slice
        fmt.Println(slice)

        i += 1
    }

    fmt.Printf("Goodbye!\n")
}

/**
Output
    oris-mbp:Golang orishadmon$ go run ./OriShadmon_slice.go
    Input Number: 6
    [6]
    Input Number: 3
    [3 6]
    Input Number: 32
    [3 6 32]
    Input Number: 0
    [0 3 6 32]
    Input Number: 7
    [0 3 6 7 32]
    Input Number: 9
    [0 3 6 7 9 32]
    Input Number: 17
    [0 3 6 7 9 17 32]
    Input Number: 1
    [0 1 3 6 7 9 17 32]
    Input Number: 5
    [0 1 3 5 6 7 9 17 32]
    Input Number: 2
    [0 1 2 3 5 6 7 9 17 32]
    Input Number: 4
    [0 1 2 3 4 5 6 7 9 17 32]
    Input Number: X
    Goodbye!
*/