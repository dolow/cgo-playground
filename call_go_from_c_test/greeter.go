package main

import "C"
import (
    "fmt"
)

//export Greet
func Greet(name string, year int) {
    fmt.Printf("Hello ! %s %d years old !\n", name, year)
}

func main() {
    panic("dont call me")
}
