package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Now().Unix()
	fmt.Println(a)
	s := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println(s)
}
