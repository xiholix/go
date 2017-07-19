package main
import(
    "encoding/json"
    "fmt"
    "strconv"
)

type Foo struct{
    Number int `json:"number"`
    Title string `json:"title"`
}

type OneMessage struct{
    MetricName string `json:"metricName"`
    Dimensions map[string]string `json:"dimensions"`
}

func mapToJson(){
    datas := make(map[string] Foo)
    for i:=0; i<10; i++{
        datas[strconv.Itoa(i)] = Foo{Number:i, Title:"test"}
    }
    jsonString, _ := json.Marshal(datas)
    fmt.Println(string(jsonString))
}


func convertMessageToJson(){
    datas := make(map[string] string)
    for i:=0; i<10; i++{
        datas[strconv.Itoa(i)] = "test"
    }
    message := OneMessage{MetricName:"hello", Dimensions:datas}
    t := message

    f := []OneMessage{message, t}
    jsonString, _ := json.Marshal(f)
    fmt.Println(string(jsonString) )
}

func main(){
    convertMessageToJson()
}