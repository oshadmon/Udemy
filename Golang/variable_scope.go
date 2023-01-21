package main
import "fmt"

var x int = 4

func f(){
    /**
    call variable from outside
    */
    fmt.Printf("%d\n", x)
}

func g(){
    /**
    call variable from within the function
    */
    var y string = "a"
    fmt.Printf("%s\n", y)
}

func main(){
    f()
    g()
}