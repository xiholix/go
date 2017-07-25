package main
import(
    "fmt"
)

func main(){
   a := []string{}
   for _, b := range a{
       fmt.Println(b)
   }

   fmt.Println("hello world")
   b := map[string]interface{}{"a":1}
   fmt.Println(b["b"])

}