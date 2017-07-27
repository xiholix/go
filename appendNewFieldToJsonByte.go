package main
import(
    "fmt"
    "encoding/json"
)

func main(){
    a := map[string]string{"one":"hello", "two":"world"}
    b, _ := json.Marshal(a)
    var v  map[string]interface{}
    json.Unmarshal(b, &v)
    // s := v.(map[string] interface{})
    v["thrid"] = 2
    fmt.Println(v)
    t, _ := json.Marshal(v)
    fmt.Println(string(t))
}