package main

import "fmt"
import "net"
func main(){
    conn, err := net.Dial("ip4", "www.baidu.com")
    if err!=nil{
    	fmt.Println(conn)
    	fmt.Println(err)
    }else{
    	fmt.Println("not connected!")
    }

}