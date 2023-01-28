package main

import (
    "fmt"
)

// type struct Person (
//     name string
//     addr string
//     phone string
// )

func main(){
    var p1 Person
    p1.name = "Ori"

    barr, err := json.Marshal(p1)
    fmt.Println(barr)
}