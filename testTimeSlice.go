package main 

import(
  "fmt"
  "time"
)


func StartCac(){
    t1 := time.Now()
    for i :=0; i<10000; i++{
    	fmt.Print("*")
    }
    elapsed := time.Since(t1)
    fmt.Println("app elapsed: ", elapsed)
}

func main(){
	StartCac()
}