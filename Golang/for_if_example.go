package main
import "fmt"

func switch_case(x int) {
    switch x {
        case 1: fmt.Printf("Case 1\n")
        case 2: fmt.Printf("Case 2\n")
        default: fmt.Printf("Default\n")
    }
}

func main(){
    // for / if/else
    for i := 0 ; i < 10 ; i++ {
        if i % 2 == 0 {
            fmt.Printf("%d - even\n", i)
        } else {
            fmt.Printf("%d - odd\n", i)
        }
    }


    switch_case(1)
    switch_case(2)
    switch_case(0)

}