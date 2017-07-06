package main 
import(
    "log"
    "os"
)

func main(){
     fileName := "InfoFirst.log"
     logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)

     defer logFile.Close()
     if err!=nil{
     	log.Fatalln("open file error")
     }

     debugLog := log.New(logFile, "[Info]", log.Llongfile|log.Ldate|log.Ltime)
     debugLog.Println("A info message here")
     debugLog.SetPrefix("[Debug]")
     debugLog.Println("A Debug Message here")

}