package main 

import(
    "fmt"
    "encoding/json"
)


type Message struct{
    Name string
    Age int
    Address string
}



func main(){
   m := Message{Name:"xie", Age:23, Address:"hengyang"}
   b, _ := json.Marshal(m)
   var k Message
   json.Unmarshal(b, &k)
   fmt.Println(m)
   fmt.Println(b)
   fmt.Println(k)
   fmt.Println(k.Name)

   var a=10
   fmt.Println(a)
   var binary= Message{Age:a}
   fmt.Println(binary)
   newBinary, _ := json.Marshal(binary)
   fmt.Println(newBinary)

}