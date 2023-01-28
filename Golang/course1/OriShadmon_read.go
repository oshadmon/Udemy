package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Name struct {
    fname string
    lname string
}

func __get_file() string {
    /**
    User input to get file name
    :params:
        fileName:string - user defined file name
    :return:
        fileName
    */
    var fileName string

    fmt.Printf("File Name: ")
    fmt.Scan(&fileName)

    return fileName
}


func main() {
    /**
    Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a
    text file which contains a series of names. Each line of the text file has a first name and a last name, in that order,
    separated by a single space on the line.

    Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
    Each field will be a string of size 20 (characters).

    Your program should prompt the user for the name of the text file. Your program will successively read each line of
    the text file and create a struct which contains the first and last names found in the file. Each struct created will
    be added to a slice, and after all lines have been read from the file, your program will have a slice containing one
    struct for each line in the file. After reading all lines from the file, your program should iterate through your
    slice of structs and print the first and last names found in each struct.

    :params:
        fileName:string - user defined file
        names:array - array of names
        file, err - content from file as bytes || error message
    */

    var fileName string = __get_file()
    var names []Name

    file, err := os.Open(fileName) // read file
    if err != nil {
        fmt.Printf("Failed to open file %s (Error: %s)\n", fileName, err)
        os.Exit(1)
    }
    defer file.Close() // close file

    scanner := bufio.NewScanner(file) // store content from file to be used
    for scanner.Scan() {
        line := scanner.Text()
        firstName, lastName := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
        names = append(names, Name{fname: firstName, lname: lastName})
    }

    for _, name := range names {
        fmt.Printf("First Name: %s | Last Name: %s\n", name.fname, name.lname )
    }
}

/**
Sample output:
oris-mbp:Golang orishadmon$ go run OriShadmon_read.go
File Name: names.txt
First Name: Michael | Last Name: Smith
First Name: Emily | Last Name: Johnson
First Name: Matthew | Last Name: Williams
First Name: Olivia | Last Name: Jones
First Name: David | Last Name: Brown
First Name: Isabella | Last Name: Garcia
First Name: Joseph | Last Name: Martinez
First Name: Sophia | Last Name: Rodriguez
First Name: Jacob | Last Name: Carter
First Name: Madison | Last Name: Phillips
First Name: Ryan | Last Name: Evans
First Name: Chloe | Last Name: Turner
First Name: Nathan | Last Name: Lewis
First Name: Elizabeth | Last Name: Parker
First Name: Aaron | Last Name: Nelson
First Name: Natalie | Last Name: Carter
First Name: Adam | Last Name: Mitchell
First Name: Samantha | Last Name: Perez
First Name: Anthony | Last Name: Roberts
First Name: Gabriella | Last Name: Cook
First Name: Joshua | Last Name: Cooper
First Name: Audrey | Last Name: Meyer
*/