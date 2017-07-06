package main
import(
    "fmt"
    "time"
)


type Data struct{
    Data_channel chan int
}


func(this *Data)Send(){
    for{
        this.Data_channel <- 1
        time.Sleep( 2*time.Second)
    }
}


func(this *Data)Receive(){
    for{
        d := <- this.Data_channel
        fmt.Println(d)
    }
}


func main(){
    data_chan := make(chan int, 4)
    send := Data{Data_channel:data_chan}
    receive := Data{Data_channel:data_chan}

    go send.Send()
    time.Sleep(6*time.Second)
    go receive.Receive()
    for{
        
    }
}