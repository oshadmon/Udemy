package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "math/rand"
)

func userValue() [10]int {
    /**
    Request user to insert 10 random values - if not set, will randomly select a value between 0 and 10
    :params:
        x:list - list of 10 integer values
        inputReader - user input
    :return:
        x
    **/
    var x [10]int
    for i := 0; i < len(x); i++ {
        fmt.Printf("Enter value %d: ", i + 1)
        inputReader := bufio.NewReader(os.Stdin)
        input, _ := inputReader.ReadString('\n')
        intInput, err := strconv.Atoi(strings.TrimSpace(input))
        if err != nil {
            x[i] = rand.Intn(11)
        } else {
            x[i] = intInput
        }
    }
    return x

}

func BubbleSort(x [10]int) [10]int {
    for i:=0 ; i < len(x) ; i++ {
        for j := 0 ; j < len(x)-i-1 ; j++ {
            if x[j] > x[j+1] {
                tmp := x[j]
                x[j] = x[j+1]
                x[j+1] = tmp
            }
        }
    }
    return x
}

func main() {
    /**
    Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers.
    The program should print the integers out on one line, in sorted order, from least to greatest. Use your favorite
    search tool to find a description of how the bubble sort algorithm works.
    :params:
        x:list - list of generated values 
        sorted_x:list - list of sorted values
    **/
    x := userValue()
    sorted_x := BubbleSort(x)
    fmt.Printf("%v\n", sorted_x)
}

/**
Results:
$ go run Golang/course2/OriShadmon_BubbleSort.go
Enter value 1: 4
Enter value 2: 5
Enter value 3: 2
Enter value 4: 1
Enter value 5: 0
Enter value 6: 8
Enter value 7: 7
Enter value 8: 6
Enter value 9: 9
Enter value 10: 3
[0 1 2 3 4 5 6 7 8 9]
*/