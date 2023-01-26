/**
useful tool: https://golangtutorial.dev/tips/foreach-loop-go/#:~:text=There%E2%80%99s%20no%20foreach%20loop%20in%20Golang.%20However%20we,variables%20if%20present%20and%20then%20executes%20the%20block.
*/
package main
import "fmt"

func main(){
    // declare array
    var x [5]int
    // preset values in array size 4
    y := [4]int{1, 2, 3, 4}

    x[0] = 2 // set value

    // print array x value @ 0
    fmt.Println(x[0])

    // for loop -- i(ndex) | v(alue)
    for i, v := range y {
        fmt.Printf("index: %d | value: %v\n", i, v)
    }

    // for loop print only value
    for _, v := range y {
        fmt.Println(v)
    }
}