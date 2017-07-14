package main

import(
    "os"
    "fmt"
    "net"
    "encoding/json"
)

type Message struct{
    Srcid int
    Objid int
    Time string
    Msg string
}

func checkError(err error){
    if  err != nil {
        fmt.Println("Error: %s", err.Error())
        os.Exit(1)  
    }
}

func recvUDPMsg(conn *net.UDPConn){
    for{
        var buf [200]byte
    
        n, raddr, err := conn.ReadFromUDP(buf[0:])
        if err != nil {
            return  
        }
    t := Message{}
    json.Unmarshal(buf[0:], &t)
    fmt.Printf("%+v\n", t)
    fmt.Println("msg is ", string(buf[0:n]))

    //WriteToUDP
    //func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error) 
    _, err = conn.WriteToUDP([]byte("nice to see u"), raddr)
    checkError(err)
    }
    
}

func main() {
    udp_addr, err := net.ResolveUDPAddr("udp", ":11110")
    checkError(err) 
    conn, err := net.ListenUDP("udp", udp_addr)
    defer conn.Close()
    checkError(err) 
    fmt.Printf("hello world")

    //go recvUDPMsg(conn)
    recvUDPMsg(conn)
}