package main
import(
    "fmt"
)

func main(){
    var user_input float32

    fmt.Printf("Floating Point Number? ")
    _, err := fmt.Scan(&user_input)
    if err != nil {
        fmt.Println("Enter a valid number")
    } else{
        var converted_value int = int(user_input)
        print("Converted value: ", converted_value, "\n")
    }


}