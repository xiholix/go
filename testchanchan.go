package main 
import(
	"fmt"
)


func test(values chan chan int){
	for v := range values{
		fmt.Println(v)   //v 变为chan int类型
		a := <-v
		fmt.Println(a)
	}
	fmt.Println("end")
}


func main(){
	a := make(chan chan int)
	b := make(chan int)
	fmt.Println(a)
	fmt.Println(b)
	go func(){b<-1}()
	go func(){a<-b}()
	test(a)
}