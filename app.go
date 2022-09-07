package main

import (
    "fmt"
    "time"
)

func main() {
    for {
        fmt.Printf("hello world - curse of the demo?\n")

        time.Sleep(5 * time.Second)
    }
}
