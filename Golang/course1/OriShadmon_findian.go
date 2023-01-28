package main

import (
    "bufio"
    "fmt"
    "os"
)

func checkString(user_input string) bool {
    /**
    Check whether values a, i or n is found in a given string
    :args:
        user_input:string - user input to check
    :params:
        status:bool
    :return:
        if a, i or n is found returns True
        else returns False
    */
    var status bool = false
    for _, char_variable := range user_input {
        if (char_variable == 'A' || char_variable == 'I' || char_variable == 'N' ||
            char_variable == 'a' || char_variable == 'i' || char_variable == 'n'){
            status = true
        }
    }

    return status
}

func main() {
    /**
        Write a program which prompts the user to enter a string. The program searches through the entered string
        for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with
        the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print
        “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters
        are upper-case or lower-case.
    */

    // user input
    inputReader := bufio.NewReader(os.Stdin)
    fmt.Printf("User Input: ")
    user_input, _ := inputReader.ReadString('\n')

    // check if a, i or n was found
    if checkString(user_input) == true {
        fmt.Println("Found!")
    } else {
        fmt.Println("Not Found!")
    }

}
