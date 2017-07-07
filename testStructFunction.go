package main 
import(
    "fmt"
)


type Message struct{
    Age int
    Name string
}

func(this *Message)Test( i int){
    fmt.Println(this.Age)
    fmt.Println(i)
}

func main(){
    m := Message{Age:10, Name:"xie"}
    m.Test(3)
    // Test(&m, 4)   // 这种调用方式不正确
}