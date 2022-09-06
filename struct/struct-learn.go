package main
 
import (
    "fmt"
)

type(

    Learns struct{
        timepass int16
        pratice int32
    }

    Goals struct{
        name string
        Content string
    }

    Peoples struct {
        name  string
        child *Peoples
    }
)


func main() {

    learn := &Learns{pratice:16}
    println(learn)
    fmt.Println(learn)
    println(learn.pratice)

    goal := new(Goals)
    goal.name = "rpg"
    goal.Content = "code a place to go"
    fmt.Println(goal)

    var study Goals
    study.name = "find"
    study.Content = "things to do"
    fmt.Println(study)

    
    relation := &Peoples{name:"爷爷",child:&Peoples{name:"爸爸",child:&Peoples{name:"我"}}}
    fmt.Println(relation)
    println(relation.child.child.name)

}
