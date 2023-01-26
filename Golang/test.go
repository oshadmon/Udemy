package main

import (
    "flag"
    "fmt"
)

func main() {
    var name string

    flag.StringVar(&name, "name", "John Doe", "Your name")

    var age int
    flag.IntVar(&age, "age", 25, "Your age")

    flag.Parse()

    fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}
