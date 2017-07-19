package main
import(
    "fmt"
    "encoding/json"
)

type MacSig struct{
    Mac int
    Sig int
}

type OneMessage struct{
    Src_ip string 
    Device_id int
    time int
    AirStations []MacSig
}


func main(){
    msg := []byte(`{"src_ip": "100.116.185.114", "airstations": [{"mac": "0456040790D8", "sig": 47}, {"mac": "F48E92878CFF", "sig": 37}, {"mac": "9801A7403137", "sig": 31}, {"mac": "04C23E3E3539", "sig": 56}, {"mac": "683E34D0B9BD", "sig": 27}, {"mac": "440444318045", "sig": 30}, {"mac": "20F41B068C0D", "sig": 47}, {"mac": "E8088B34E9B2", "sig": 30}], "device_id": 80162, "time": 1500353223}`)
    var s OneMessage
    json.Unmarshal(msg, &s)
    fmt.Println(s)
    fmt.Printf("%v", s)
}