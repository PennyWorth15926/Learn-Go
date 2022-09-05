package main

import (
    "fmt"
)

func main() {

    // 声明一个二维整型数组，两个维度的长度分别是 4 和 2
    var array [4][2]int
    // 使用数组字面量来声明并初始化一个二维整型数组
    array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
    // 声明并初始化数组中索引为 1 和 3 的元素
    array = [4][2]int{1: {20, 21}, 3: {40, 41}}
    // 声明并初始化数组中指定的元素
    array = [4][2]int{1: {0: 20}, 3: {1: 41}}  

    fmt.Println(array)

    // 声明一个 2×2 的二维整型数组
    var array1 [2][2]int
    // 设置每个元素的整型值
    array1[0][0] = 10
    array1[0][1] = 20
    array1[1][0] = 30
    array1[1][1] = 40

    fmt.Println(array1)

    var array2 [2][2]int

    array2 = array1

    fmt.Println(array2)

}