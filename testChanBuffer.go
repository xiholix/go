package main 
import(
    "fmt"
    "time"
)


func produce_message(chanBuffer chan int){
	for{
		chanBuffer <- 1
		time.Sleep(time.Second * 1)
	   }
}


func main(){
	chanBuffer := make(chan int, 100)
    go produce_message(chanBuffer)
    
    time.Sleep(time.Second*5)
    for i:=0; i<10; i++{
    	v := <- chanBuffer
    	fmt.Println(v)
    }

}