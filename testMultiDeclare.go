package main
import(
	"fmt"
)


func main(){
	a := 10
	fmt.Println(a)
	a, b:= "hel", "test"  // xhl cannot use "hell" as type int in assignment
	fmt.Println(a)
	fmt.Println(b)
}