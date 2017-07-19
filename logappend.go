package main
import(
    "os"
    "log"
    "fmt"
)

func appendToLog(){
    path := "appendFile.txt"
    f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    if err!=nil{
        fmt.Println(err)
    }
    logger := log.New(f, "Logger:", log.Llongfile|log.LstdFlags)
    // logger.SetFlags(1)
    logger.SetPrefix("Error")
    logger.Print("Hello, log file!\n\nhello")
    // logger.Output(4, "hello world")

}


func SysLog(){

}

func main(){
   appendToLog()
}