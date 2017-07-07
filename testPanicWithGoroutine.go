package main 
import(
    "fmt"
    "time"
    "os"
)


func OneGoroutine(){
    fmt.Println("in go routine")
    time.Sleep(3*time.Second)
    panic("error has occured!!!")
}


// func main(){   //不能在main函数中用recover中断goroutine的panic
//     defer func(){
//         if r:=recover(); r !=nil{
//             fmt.Println("catch error ")
//             fmt.Println("the recover is %v", r)
//         }
//     }()
//     go OneGoroutine()

//     for{

//     }
// }



func OneGoroutineWithErrorChannel(err chan<- error){
    fmt.Println("in go routine")
    time.Sleep(3*time.Second)
    err <- fmt.Errorf("hello world")
}


func main(){
    err := make(chan error)
    fmt.Println("begin goroutine")
    go OneGoroutineWithErrorChannel(err)
    for{
        select{
            case e :=<-err:
                 fmt.Println(e)
                 os.Exit(99)
            default:
        }
    }

}