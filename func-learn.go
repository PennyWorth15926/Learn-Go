package main
 
import (
    "fmt"
)

type(    

    People struct {
        name  string
        goal *Goal
    }

    Goal struct{
        name string
        Content string
    }
)


func Changes(whisper *Goal, msg []byte) (*People, error) {
    person := new(People)
    person.goal = whisper
    println(msg)
    fmt.Println(msg)
    return person, nil
}  

func main() {

    whisper := &Goal{Content:"sleep a little"}
    bytemsg := []byte("Hl Asong!")
    person,err := Changes(whisper,bytemsg)
    if err != nil {
        fmt.Println("error message", err)
    }
    fmt.Println(person)
    fmt.Println(person.goal)

}


