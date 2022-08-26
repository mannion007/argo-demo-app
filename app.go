package main

import (
    "fmt"
    "time"
)

func main() {
    for {
        fmt.Printf("hello world!!! This is another update\n")

        time.Sleep(5 * time.Second)
    }
}
