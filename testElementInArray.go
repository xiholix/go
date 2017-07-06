package main 
import(
	"fmt"
)

var ValidActionKeys = map[string] int{"RESETPOWER":0, "RESETPASSWD":1}

func main(){
     _, ok := ValidActionKeys["RESETPOWER"]
	 fmt.Println(ok)
     _, ok = ValidActionKeys["he"]
	 fmt.Println(ok)
}