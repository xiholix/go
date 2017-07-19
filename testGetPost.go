package main
import(
    "net/http"
    "fmt"
    "io/ioutil"
    "strings"
    "net/url"
)


func Http_get(){
    url := "http://www.01happy.com/demo/accept.php?id=1"
    resp, err:= http.Get(url)
    if err!=nil{
        fmt.Println(err)
    }
    defer resp.Body.Close()
    body, err:=ioutil.ReadAll(resp.Body)
    if err!=nil{
        fmt.Println(err)
    }

    fmt.Println(string(body))
}


func Http_post(){
    url := "http://www.01happy.com/demo/accept.php"
    resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader("name=cjb,test=12"))
    //若x-www-form-urlencoded的拼写不正确，并不会报错，但是无法正确post数据
    if err!=nil{
        fmt.Println(err)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err!=nil{
        fmt.Println(err)
    }

    fmt.Println(string(body))
}

func HttpPostForm(){
    resp, err := http.PostForm("http://www.01happy.com/demo/accept.php",
        url.Values{"key": {"Value"}, "id": {"123"}})
 
    if err != nil {
        // handle error
    }
 
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }
 
    fmt.Println(string(body))

}

func main(){
   Http_post()
}

