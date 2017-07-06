package main 
import(
    "fmt"
    "time"
)

var chan_int chan int = make(chan int, 5)
var count = 0

func Producer(){
    for{
        time.Sleep(3*time.Second)
        chan_int <- 1
    }
}


func consumer(){
    for{
        select{
            case <-chan_int:
                fmt.Println(count)
                count = 0
                fmt.Println("consumer")
            default:
                // fmt.Println("default")
                count += 1
        }
    }
}


func NewProducer(){
    for{
        select{
            case chan_int <- 1:
                fmt.Println(count)
                count = 0
            default:
                count += 1
        }
    }
}


func NewConsumer(){
    for{
         <- chan_int
        time.Sleep(3*time.Second)
    }
}

func main(){
    go NewProducer()
    go NewConsumer()
    for{
        time.Sleep(3*time.Second)
    }
}