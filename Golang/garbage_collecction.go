package main
import "fmt"


func foo() *int {
    var x int = 1
    return &x // returns address of x
}

func main() {
    var y *int
    y = foo()
    fmt.Printf("%d\n", y) // print address of result from x
    fmt.Printf("%d\n", *y) //
}