package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "sort"
    "strconv"
    "strings"
    "sync"
//     "time"
)

var wg sync.WaitGroup

// func generateList(listSize int) []int {
//     /**
//     Generate a list of (random) integers with a given list size
//     :args:
//         listSize:int - list size
//     :params:
//         list:list - generated list of values
//     :return:
//         list
//     */
//     rand.Seed(time.Now().UnixNano())
//     list := make([]int, listSize)
//     var i int = 0
//     for i < listSize {
//         value := rand.Intn(listSize * 2)
//         status := false
//         for _, item := range list { // check if integer in the list
//             if value == item {
//                 status = true
//                 break
//             }
//         }
//         if !status { // insert value to list
//             list[i] = value
//             i += 1
//         }
//     }
//     return list
// }

func userInput() []int {
    /**
    Allow user input
    :params:
        numbers:list - list of numbers
    :return:
        numbers
    */
    var numbers []int

    fmt.Printf("Enter Array Size: ")
    inputReader := bufio.NewReader(os.Stdin)
    input, _ := inputReader.ReadString('\n')
    intInput, err := strconv.Atoi(strings.TrimSpace(input))
    if err == nil {
        numbers = make([]int, intInput)
    } else {
        numbers = make([]int, 16)
    }

    for i:=0 ; i < len(numbers) ; i++ {
        fmt.Printf("Enter Number: ")
        inputReader := bufio.NewReader(os.Stdin)
        input, _ := inputReader.ReadString('\n')
        intInput, err := strconv.Atoi(strings.TrimSpace(input))
        if err == nil {
            numbers[i] = intInput
        } else{
            numbers[i] = rand.Intn(100)
        }
    }

    return numbers

}


func partition(numbers []int, numParts int, partSize int) [][]int {
    /**
    Partition the data
    :args:
        numbers:list - list of numbers
        numParts:int - number of parts
        partSize:int - size of each partition
    :params:
        parts:list - partitioning of numbers
    :return:
        parts
    **/
    parts := make([][]int, numParts)

    for i := 0; i < numParts; i++ {
        start := i * partSize
        end := start + partSize
        if i == numParts-1 {
            end = len(numbers)
        }
        parts[i] = numbers[start:end]
    }
    return parts
}

func sorting(group_id int, part []int, wgIns *sync.WaitGroup) {
    /**
    function call to sort list
    */
    fmt.Println("Group ", group_id, "pre-sorting: ", part)
    sort.Ints(part)
    fmt.Println("Group ", group_id, "post-sorting: ", part)
    wgIns.Done()
}

func main() {
    /**
    Main
    :params:
        numbers:list - list of generated values
        numParts:int - how many partitions
        partSize:int - how many values per partition
    :params:
        parts:list - a list of lists with subset of values
        sortedParts:list - sorted list
    */
//     var numbers []int = generateList(16)
    var numbers []int = userInput()
    var numParts int = 4
    var partSize int = len(numbers)/numParts
    var sortedParts []int

    parts := partition(numbers, numParts, partSize)

    for i := 0; i < numParts; i++ { // sort per partition
        wg.Add(1)
        go sorting(i+1, parts[i], &wg)
    }
    wg.Wait()

    for i := 0; i < numParts; i++ { // merge lists into one
        sortedParts = append(sortedParts, parts[i]...)
    }

    sort.Ints(sortedParts)
    fmt.Println("Unsorted List: ", numbers)
    fmt.Println("Sorted List: ", sortedParts)
}

/**
go run Golang/course3/OriShadmon_Partition.go
Enter Array Size: 16
Enter Number: 8
Enter Number: 13
Enter Number: 12
Enter Number: 40
Enter Number: 5
Enter Number: -6
Enter Number: 6
Enter Number: 0
Enter Number: 1212
Enter Number: 164
Enter Number: 22
Enter Number: 27
Enter Number: 75
Enter Number: 44
Enter Number: 0
Enter Number: 3
Group  1 pre-sorting:  [8 13 12 40]
Group  1 post-sorting:  [8 12 13 40]
Group  2 pre-sorting:  [5 -6 6 0]
Group  3 pre-sorting:  [1212 164 22 27]
Group  3 post-sorting:  [22 27 164 1212]
Group  4 pre-sorting:  [75 44 0 3]
Group  4 post-sorting:  [0 3 44 75]
Group  2 post-sorting:  [-6 0 5 6]
Unsorted List:  [8 12 13 40 -6 0 5 6 22 27 164 1212 0 3 44 75]
Sorted List:  [-6 0 0 3 5 6 8 12 13 22 27 40 44 75 164 1212]
*/