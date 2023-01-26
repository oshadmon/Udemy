// code is based on ChatGPT
package main

import (
    "fmt"
)

type HashTable struct {
    size int
    slots map[int]string
}

func (ht *HashTable) Insert(key int, value string) {
    ht.slots[key] = value
}

func (ht *HashTable) Lookup(key int) string {
    return ht.slots[key]
}

func main() {
    ht := &HashTable{size: 11, slots: make(map[int]string)}
    ht.Insert(1, "value1")
    ht.Insert(2, "value2")
    fmt.Println(ht)
    fmt.Println(ht.Lookup(1)) // Output: "value1"
    fmt.Println(ht.Lookup(2)) // Output: "value2"
}
