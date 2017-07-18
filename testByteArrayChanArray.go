package main
import(
    "fmt"
    "time"
)





func consume(t chan []byte){
    msg:= <- t
    fmt.Println(string(msg))
}

func main(){
    var multiSend = make([] chan []byte, 6) ;
   multiSend[0] = make(chan []byte)
   go consume(multiSend[0])
   time.Sleep(3*time.Second)
   multiSend[0]<- []byte("hello world")
   for{
       time.Sleep(time.Second)
   }
}
