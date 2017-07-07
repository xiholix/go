package main
import(
    "fmt"
)


type MyStruct struct{
    Age int
    Name string
}

func NewMyStruct(age int, name string) *MyStruct{  // xhl书上说的构造函数，应该是通过他模拟构造函数的作用但是还是需要显示的调用
    return &MyStruct{Age:age, Name:name}
}


func main(){
    t := new(MyStruct)
    fmt.Println(t)
}