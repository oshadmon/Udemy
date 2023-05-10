package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Animal interface {
    Eat() string
    Move() string
    Speak() string
}

var animals = map[string]Animal{
    "cow":   Cow{},
    "bird":  Bird{},
    "snake": Snake{},
}


func user_input() (string, string, string) {
    /**
    Code to accept user input and extrapolate animal & command
    :retun:
        animal, command
    */
    var option string = ""
    var param1 string = ""
    var param2 string = ""

    fmt.Print("> ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := strings.ToLower(scanner.Text())
    command := strings.Fields(input)

    if len(command) != 3 {
        if command[0] == "exit" {
            fmt.Printf("Goodbye!")
            os.Exit(0)
        }
        fmt.Println("Invalid input")

    } else {
        option = command[0]
        param1 = command[1]
        param2 = command[2]
    }
    return option, param1, param2
}


type Cow struct {
    // declare a cow
    name string
}

func (c Cow) Eat() string {
    return "grass"
}

func (c Cow) Move() string {
    return "walk"
}

func (c Cow) Speak() string {
    return "moo"
}

type Bird struct {
    name string
}

func (b Bird) Eat()  string {
    return "worms"
}

func (b Bird) Move() string {
    return "fly"
}

func (b Bird) Speak() string {
    return "peep"
}

type Snake struct {
    name string
}

func (s Snake) Eat() string {
    return "mice"
}

func (s Snake) Move() string {
    return "slither"
}

func (s Snake) Speak() string {
    return "hsss"
}

func main() {
    for {
        option, param1, param2 := user_input()

        switch option {
        case "newanimal":
            name := param1
            switch param2 {
            case "cow":
                animals[name] = Cow{name}
            case "bird":
                animals[name] = Bird{name}
            case "snake":
                animals[name] = Snake{name}
            default:
                fmt.Println("Invalid animal type")
                continue
            }
            fmt.Println("Created it!")
        case "query":
            name := param1
            if a, ok := animals[name]; ok {
                switch param2 {
                case "eat":
                    fmt.Printf("A %s eats %s\n", name, a.Eat())
                case "move":
                    fmt.Printf("A %s moves %s\n", name, a.Move())
                case "speak":
                    fmt.Printf("A %s sounds %s\n", name, a.Speak())
                default:
                    fmt.Printf("Unknown command: %s\n", param2)
                }
            } else {
                fmt.Println("Animal not found")
                continue
            }
        default:
            fmt.Println("Invalid command")
            continue
        }
    }
}

/**
> query cow move
A cow moves walk
> query bird speak
A bird sounds peep
> query snake eat
A snake eats mice
> addanimal python snake
Invalid command
> newanimal python snake
Created it!
> query python eat
A python eats mice
> query python speak
A python sounds hsss
> query pthon move
Animal not found
> qury python move
Invalid command
> query python move
A python moves slither
> exit
Goodbye!
*/