package main 
import(
	"fmt"
	"reflect"
	"encoding/json"
)

type Person struct{
    Name string
	Age  int
}


func convert_struct_to_map(obj interface{}) map[string] interface{}{
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string] interface{} )
	for i:=0; i < t.NumField(); i++{
		data[t.Field(i).Name ] = v.Field(i).Interface()
	}
	return data
}

func main(){
    a := Person{Name:"xie", Age:23}
	fmt.Println(a)

	t := convert_struct_to_map(a)
	fmt.Println(t)
	s, _ := json.Marshal(a)
	s2, _ := json.Marshal(t)
    fmt.Println(string(s))
	fmt.Println(string(s2))
}