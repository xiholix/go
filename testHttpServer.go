package main
import(
    "io"
    "net/http"
    "log"
    "time"
    "strconv"
    "fmt"
    "strings"
    "encoding/json"
    "os"
)

var test string = "hello"

func HelloServer(w http.ResponseWriter, req *http.Request){
    // fmt.Println(req)
    // fmt.Println(req.URL.Query())
    // t:= req.URL.Query()
    // if len(t)>0{
    //    time.Sleep(time.Second*10)
    // }
        
    // fmt.Println(t)
	io.WriteString(w, "hello world!\n")
    req.ParseForm()
    if (len(req.Form["id"])>0){
        io.WriteString(w, req.Form["id"][1])
    }
    // realIp := req.Header.Get("X-Forwarded-For")
    realIp := req.Header.Get("User-Agent")
    io.WriteString(w, realIp+"\n")
    io.WriteString(w, strconv.FormatInt(time.Now().Unix(), 10) )
    // var a int = time.Now().Unix()
    // fmt.Println(a)
    fmt.Println(test)
}


func check_params_valid(w http.ResponseWriter, req *http.Request) bool{
    if( len(req.Form["ip"])==0 ){
        fmt.Println("Error")
        io.WriteString(w, "Error")
        return false
    }
    io.WriteString(w, "ip is correct")
    return true
}


func Test_get_param(w http.ResponseWriter, req *http.Request){
    req.ParseForm()
    val := check_params_valid(w, req)
    if val==false{
        return
    }
    fmt.Println("hello")
    didStr, _ := strconv.ParseInt(req.Form["did"][0], 10,  64)
    var s string = req.Form["ip"][0]
    t := strings.Split(s, ",")
    fmt.Println(t)
    fmt.Println(didStr)
    data := []byte("hello world")
    w.Write(data)
   
}

type Message struct{
    Name string
    Age  int
}

func Test_write_struct(w http.ResponseWriter, req *http.Request){
     a :=[]Message{ Message{Name:"xiehong", Age:22}, Message{Name:"liang", Age:23} }
     b, _ := json.Marshal(a)
     w.Write(b)
     
}


func asyn_function(){
    fmt.Println("asyn will sleep")
    time.Sleep(3*time.Second)
    fmt.Println("asyn awake")
    os.Exit(99)

}

func Test_asynchronous(response http.ResponseWriter, request *http.Request){
     fmt.Println("begin sleep!")
     io.WriteString(response, "hello world")
     go asyn_function()   // xhl 可以这样使用
     fmt.Println("end of Test_asyn")
}


func main(){
     http.HandleFunc("/hello", Test_asynchronous)
     err := http.ListenAndServe(":12345", nil)
     if err!= nil{
     	log.Fatal("ListenAndServe: ", err)
     }
}