package main
import "fmt"

type Grade int
const (
    A Grade = iota
    B
    C
    D
    F
)


func main(){
    fmt.Printf("A - %d | B - %d | C - %d | D - %d | F - %d\n", A, B, C, D, F)


}