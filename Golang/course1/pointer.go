package main
import (
    "fmt"
)

func main() {
    var x int = 1
    var y int
    var ip *int  // ip is pointer to an int

    fmt.Print(x) // print
    fmt.Printf("\n")
    fmt.Print(y) // print 0
    fmt.Printf("\n")
    fmt.Print(ip) // print nil
    fmt.Printf("\n")

    ip = &x
    y = *ip

    fmt.Print(y) // print 1 (data in memory)
    fmt.Printf("\n")
    fmt.Print(ip) // print address of x
    fmt.Printf("\n")


    ptr := new(int) // set address of a variable
    fmt.Print(ptr)
    fmt.Printf("\n")

    *ptr = 3 // set value 3 to the address of ptr
    fmt.Print(ptr)

}