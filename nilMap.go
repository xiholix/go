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
    if t==nil{
        fmt.Println("t is nil")
    }
}