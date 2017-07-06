package main 
import(
	"keepa/keepa"
	"errors"
	// "runtime"
	// "fmt"
)


func main(){
    keepa.Stage_info("hello world")
	// keepa.Stage_error("my error")
	keepa.Stage_error(errors.New("my error"))

	
}