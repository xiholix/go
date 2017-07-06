package main

import(
	"fmt"
	"strings"
)

func main(){
	a := [] string{"123", "345", "567"}
	b := make([]string, 0, 5)
	for _, v :=range(a){
		fmt.Println(v)
		t := fmt.Sprintf("ACT-%s", v)
		b = append(b, t)
	}
	
	for _, v :=range(b){
		fmt.Println(v)
		fmt.Println(strings.Split(v, "-")[1])
	}
	fmt.Println(len(b))
}