package main
import "fmt"

func main(){
    // create a string array
    arr := [...]string{"a", "b", "c", "d", "e", "f"}

    // set slice to values a-c
    slice1 := arr[0:3]

    // slice array c-e
    slice2 := arr[2:5]

    fmt.Printf("length: %d | capacity: %d\n", len(slice1), cap(slice1))
    fmt.Println(slice2)



    // create a slice of size 10 that's associated with an "array" of size 15
    slice3 := make([]int, 10, 15)
    fmt.Println(slice3, len(slice3))

    // increase the slice by 1 and set that value to 1000 [ same as slice.append(1000) in python ]
    slice3 = append(slice3, 1000)
    fmt.Println(slice3, len(slice3))
}