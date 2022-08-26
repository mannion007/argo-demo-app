package main

import (
    "fmt"
    "time"
)

func main() {
    for {
        fmt.Printf("hello world - team demo\n")

        time.Sleep(5 * time.Second)
    }
}
