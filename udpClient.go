package main

import (
    "os"
    "fmt"
    "net"
//  "io"
	"encoding/json"
)

type Message struct{
    Srcid int
    Objid int
    Time string
    Msg string
}

func main() {
    conn, err := net.Dial("udp", "127.0.0.1:11110")
    defer conn.Close()
    if err != nil {
        os.Exit(1)  
    }
    t := Message{Srcid:10, Objid:20, Time:"2012-12-01", Msg:"hello world!"}
    // conn.Write([]byte("Hello world!"))  
    v, _ := json.Marshal(t)
    fmt.Println(v)
    conn.Write(v)
    fmt.Println("send msg")

    var msg [20]byte
    conn.Read(msg[0:])

    fmt.Println("msg is", string(msg[0:10]))
}