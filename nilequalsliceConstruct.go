package main

import(
    "fmt"
)

func main(){
    a := []string{}
    fmt.Println(a)
    fmt.Println(a==nil)
    b := nil
    b = []string{"a", "b"}
    fmt.Println(b)
}