package main

import(
    "os"
    "fmt"
    "net"
    "encoding/json"
    // "bytes"
)

type Message struct{
    Srcid int  `json:"srcid"`
    Objid int  `json:"objid"`
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
    // for{
        var buf [200]byte  // xhl 此处需要固定长度，否则无法解析成功
    
        n, raddr, err := conn.ReadFromUDP(buf[0:])
        if err != nil {
            return  
        }
    t := Message{}
    // b := buf[0:]
    json.Unmarshal(buf[0:n], &t) // xhl 此处必须指定n，否则无法解析成功。
    fmt.Printf("%+v\n", t)
    fmt.Println("msg is ", string(buf[0:n]))

    //WriteToUDP
    //func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error) 
    _, err = conn.WriteToUDP([]byte("nice to see u"), raddr)
    checkError(err)
    // }
    
}


func recvUDPLongMsg(conn *net.UDPConn){
    for{
        msgs := []byte{}
        var buf [4]byte
        n := 0
        var err error

        n, raddr, err := conn.ReadFromUDP(buf[0:])
        // n, raddr, err := conn.ReadFromUDP(buf[0:])
        if err != nil {
            return  
        }
        fmt.Println(string(buf[0:]))
        for n!=0{
            // msgs = append(msgs, buf[0:])
            copy(msgs, buf[0:])
            n, raddr, err = conn.ReadFromUDP(buf[0:])
            fmt.Println(string(buf[0:]))
        }
    // newMsg := msgs.Bytes()
    
    t := Message{}
    // b := buf[0:]
    json.Unmarshal(msgs, &t) // xhl 此处必须指定n，否则无法解析成功。
    fmt.Printf("\n%+v\n", t)
    fmt.Println(n)
    fmt.Println("msg is ", string(msgs))

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