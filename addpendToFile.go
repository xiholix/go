package main 
import(
    "os"
    "fmt"
)


func main(){
	path := "appendFile.txt"
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err!=nil{
        fmt.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString("hello world!\n"); err!=nil{
		fmt.Println(err)
	}
}