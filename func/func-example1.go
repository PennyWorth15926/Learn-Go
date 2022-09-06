package main

import (
    "fmt"
)

// 遍历切片的每个元素, 通过给定函数进行元素访问
func visit(list []int, f func(int)) {

    for _, v := range list {
        f(v)
    }
}

func main() {

    func(data int) {
        fmt.Println("hello", data)
    }(100)

    // 将匿名函数体保存到f()中
    f0 := func(data int) {
        fmt.Println("hello", data)
    }
    // 使用f()调用
    f0(100)

    f1 := func(v int) {
        fmt.Println(v)
    }

    // 使用匿名函数打印切片内容
    visit([]int{1, 2, 3, 4}, f1)
}


