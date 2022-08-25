package main

import (
    "fmt"
    "time"
)

func main() {
    for {
        fmt.Printf("hello world!!! This is an update!!!\n")

        time.Sleep(5 * time.Second)
    }
}