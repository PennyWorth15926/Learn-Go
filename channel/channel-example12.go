package main

import (
    "fmt"
    "time"
)

func main() {
    done := make(chan int)

    go func() {
        time.Sleep(time.Second)
        fmt.Println("C语言中文网")
        <-done
    }()

    done <- 1
}