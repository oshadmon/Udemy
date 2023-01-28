package main
import (
    "fmt"
    "net/http"
    "net" // TCP/IP networking programming
)


func main() {
    output, err := http.Get("http://www.uci.edu") // websites
    fmt.Printf("output %s | error: %s\n\n", output, err)

    net.Dial("tcp", "uci.edu:80") // TCP/IP networking programming
}