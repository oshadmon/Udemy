package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Animal struct {
    food       string
    locomotion string
    noise      string
}

var animals = map[string]Animal{
    "cow":   Animal{"grass", "walk", "moo"},
    "bird":  Animal{"worms", "fly", "peep"},
    "snake": Animal{"mice", "slither", "hsss"},
}


func user_input() (string, string) {
    /**
    Code to accept user input and extrapolate animal & command
    :retun:
        animal, command
    */
    reader := bufio.NewReader(os.Stdin)
    var animal string = ""
    var command string = ""

    fmt.Print("> ")
    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Invalid value %s (Error: %s)", err)
        return animal, command
    }

    input = strings.TrimSpace(input)
    parts := strings.Split(input, " ")
    if len(parts) != 2 {
        if parts[0] == "exit" {
            fmt.Printf("Goodbye!\n")
            os.Exit(0)
        }
        fmt.Println("Invalid input. Please enter an animal and a command separated by a space.")
        return animal, command
    } else {
        animal = parts[0]
        command = parts[1]
    }

    return animal, command
}

func (a Animal) Eat() string {
    // return what an animal eats based on animal type
    return a.food
}

func (a Animal) Move() string{
    // return how an animal moves based on animal type
    return a.locomotion
}

func (a Animal) Speak() string {
    // return what sounds an animal moves based on animal type
    return a.noise
}


func main() {

    for {
        animalName, command := user_input()

        if animalName != "" && command != "" {
            animal, ok := animals[animalName]
            if !ok {
                fmt.Printf("Unknown animal: %s\n", animalName)
                continue
            }

            switch command {
            case "eat":
                fmt.Printf("A %s eats %s\n", animalName, animal.Eat())
            case "move":
                fmt.Printf("A %s moves %s\n", animalName, animal.Move())
            case "speak":
                fmt.Printf("A %s sounds %s\n", animalName, animal.Speak())
            default:
                fmt.Printf("Unknown command: %s\n", command)
            }
        }
    }
}

/**
> cow speak
A cow sounds moo
> snake eat
A snake eats mice
> bird move
A bird moves fly
> exit
Goodbye!
*/