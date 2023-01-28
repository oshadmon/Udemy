package main
import "fmt"

func main(){
    // declare a map
    var dictionary map[string]int
    dictionary = make(map[string]int)

    dictionary["ori"] = 123
    dictionary["roy"] = 1234

    fmt.Println(dictionary)
    fmt.Println(dictionary["ori"])
    delete(dictionary, "roy")
    fmt.Println(dictionary)

    // check if value in map
    _, p := dictionary["roy"] // expect False
    fmt.Println(p)

    dictionary["roy"] = 1234
    dictionary["noa"] = 234

    for key, value := range dictionary {
        fmt.Printf("key: %s | value: %d\n", key, value)
    }


}