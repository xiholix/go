// package main 

import(
    "fmt"
    "math/rand"
    "time"
    "encoding/json"
)


func test_random(){
	for i:=0; i<10; i++{
		fmt.Println(rand.Intn(100))
	}
}


func test_random_with_time_seed(){
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0; i<10; i++{
		fmt.Println(r.Intn(100))
	}
}


type Message struct{
	T string
	Age int
}

func generate_random_bytes(){
     for i:=0; i<10; i++{
     	a := rand.Intn(1000)
     	// fmt.Println(a)
     	m := Message{Age:a, T:"hel"}
     	// fmt.Println(m)
        b, _ := json.Marshal(m)
        // fmt.Println(err)
        fmt.Println(b)
        // var t Message
      
        // fmt.Println(er)
        // fmt.Println("t is ", t)
     }
}

func main(){
    // test_random()
    // test_random_with_time_seed()
    generate_random_bytes()
}