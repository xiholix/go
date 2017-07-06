package main 
import(
	"fmt"
	"time"
)


func main(){
    for{
		gap := 3
		time.Sleep(time.Duration(gap)*1e9)
		fmt.Println("print")
	}

}