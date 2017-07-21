package main
import(
    "net/url"
    "fmt"
)

func main(){
    a := "hello&wor/ld"
    b := url.PathEscape(a)  //pathesacpe不转&
    c := url.QueryEscape(a)

    fmt.Println(b)
    fmt.Println(c)
}