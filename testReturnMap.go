package main 
import(
	"fmt"
)


func test_return_map() (map[string] string){
     a := make(map[string] string, 5)
	 a["hello"] = "world"
     a["test"] = "mytest"
	 return a
}


func main(){
	b := test_return_map()
	fmt.Println(b["hello"])
	fmt.Println(b)
}