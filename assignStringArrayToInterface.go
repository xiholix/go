package main
import(
    "fmt"
)

func AppendData(src_data interface{}, des_data interface{}) []string{
    src_data_array := src_data.([]string)
    des_data_array := des_data.([]string)
    for _, s := range src_data_array{
        des_data_array = append(des_data_array, s)
    }
    
    return des_data_array
}


func main(){
    a := []string{"hello", "world"}
    b := []string{"talk", "is", "cheap"}
    c := AppendData(a, b)
    fmt.Println(c)
}