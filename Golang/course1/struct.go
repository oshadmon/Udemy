package main
import "fmt"

type struct Person (
    name string
    addr string
    phone string
)

func set_person(name string, addr string, phone string) Person {
    var person Person
    person.name = name
    person.addr = addr
    person.phone = phone

    return person
}

func main(){
    var dictionary map[string]Person
    dictionary = make(map[string]Person)

    dictionary["person1"] = set_person("ori","920 Edgecliff", "1234")

    for key, value := range dictionary {
        fmt.Printf("key: %s | value: %d\n", key, value)
    }
}