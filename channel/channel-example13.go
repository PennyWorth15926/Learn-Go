package main

import (
    "fmt"
)

func main() {
    done := make(chan int, 12) // 带缓存通道

    go func() {
        fmt.Println("C语言中文网")
        done <- 1
    }()

    // done <- 1
    
    <-done
}