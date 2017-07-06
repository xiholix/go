package main

import "fmt"

type testStruct struct{

	a int
	b int
}


func main(){

    t := &testStruct{}
    fmt.Println(t)
    fmt.Println(t.a)

	fmt.Println("hello world!s")

	a := struct{}{}  // xhl struct{}表示类型， struct{}{}表示初始化一个struct{}类型
	fmt.Println(a)
}