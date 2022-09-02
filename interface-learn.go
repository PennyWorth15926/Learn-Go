package main
 
import (
    "fmt"
)

// Pitaya interface operate App struct
// 接口中的方法可以作用于不同结构体
type Service interface {
    Start()  
    Log(string)  
}

type GPIO struct {
    log float64
}

type Serial struct {
    log float64  
}

func (g *GPIO) Start() {
}

func (g *GPIO) Log(msg string) {
    g.log++
    fmt.Println(msg)
}

func (g *Serial) Start() {
}

func (s *Serial) Log(msg string) {
    s.log++
    fmt.Println(msg)
}
    
func main() {

    // 结构体传给接口
    var s Service = new(GPIO)
    fmt.Println(s)
    s.Start()
    s.Log("hello")
    fmt.Println(s)
    
}