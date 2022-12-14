package main

import "fmt"

func main() {

    // 创建一个整型切片，并赋值
    slice := []int{10, 20, 30, 40}
    // 迭代每个元素，并显示值和地址
    for index, value := range slice {
        fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice[index])
        }

    // 创建一个整型切片，并赋值
    slice2 := []int{10, 20, 30, 40}
    // 迭代每个元素，并显示其值
    for _, value2 := range slice2 {
        fmt.Printf("Value: %d\n", value2)
        }
}