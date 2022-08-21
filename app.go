package main

import (
    "fmt"
    "time"
)

func main() {
    for {
        fmt.Printf("hello world!\n")

        time.Sleep(5 * time.Second)
    }
}