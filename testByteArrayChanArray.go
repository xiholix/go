package main
import(
    "fmt"
    "time"
)





func consume(t chan []byte){
    msg:= <- t
    fmt.Println(string(msg))
}

type testStruct struct{
    t []chan[] byte
}

func main(){
   var multiSend = make([] chan []byte, 0, 5) ;
   t := testStruct{t:multiSend}
   multiSend = append(multiSend, make(chan []byte) )
   fmt.Println(len(multiSend))
   go consume(multiSend[0])
   time.Sleep(3*time.Second)
//    multiSend[0]<- []byte("hello world")
   t.t[0]<- []byte("hello world")

   for{
       time.Sleep(time.Second)
   }
}
