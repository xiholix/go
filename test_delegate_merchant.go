package main
import(
    "fmt"
    "dinfo/dinfo"
)

func main(){
    t := dinfo.AdwallBasicConfig(10, "hello")
    fmt.Println(t)
    fmt.Println(t["error_code"])
}