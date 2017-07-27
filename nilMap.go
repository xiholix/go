package main

import (
    "fmt"
)


func result() map[string] interface{}{
    return nil
}


func main(){
    t := result()
    fmt.Println(t["hello"])
    t["hello"] = "tttttt"
    if t==nil{
        fmt.Println("t is nil")
    }
    fmt.Println(t)
}