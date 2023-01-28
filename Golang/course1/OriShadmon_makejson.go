package main
import (
    "bufio"
    "fmt"
    "os"
    "encoding/json"
)

type Person struct {
    Name string `json:"name"`
    Addr string    `json:"addr"`
}


func __get_answer(question string) string {
    /**
    Store user provided answer (multiple words) into a single variable
    :args:
        question:str - question to ask
    :params:
        scanner - Stdin user input
    :return:
        returns content provided by user - if not provided returns ""
    */

    fmt.Printf(question)
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        return scanner.Text()
    }
    if err := scanner.Err(); err != nil {
        return ""
    }
    return ""
}

func main(){
    /**
    based on user input create a JSON
    :global:
        Person:struc - structure with person information
    :param:
        name:str - user inputted name
        addr:str - user inputted address
        person:struc  - Person structure with new values
        jsonData:dict - person structure converted to JSON
    */
    var name string = __get_answer("Name: ")
    var addr string = __get_answer("Address: ")

    person := Person{Name: name, Addr: addr}

    jsonData, _ := json.Marshal(person)

    fmt.Println(string(jsonData))

}

/**
Sample Output:
oris-mbp:Golang orishadmon$ go run OriShadmon_makejson.go
Name: Ori
Address: a street
{"name":"Ori","addr":"a street"}
*/