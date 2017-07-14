package main
import(
    "fmt"
    "time"
)


func main(){
    var gmtLoc = time.FixedZone("GMT", 0)
    t := time.Now().In(gmtLoc).Format(time.RFC1123)
    fmt.Println("%+v\n", t)   //xhl 此处必须有\n,否则会打印%+v出来
    fmt.Println("%+v\n", gmtLoc)
}