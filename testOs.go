package main 
import(
	"fmt"
	"os"
    "os/exec"
    "strconv"
)


func test(){
	fmt.Println("function ")
	os.Exit(1)
	fmt.Println("after exit")
}


func getPid(){
    i := os.Getpid()
    fmt.Println(i)
}

func main(){
    i := os.Getpid()
    command_str := "echo "+ strconv.Itoa(i)+ " > this.pid"
    fmt.Println(command_str)
    o, err := exec.Command("bash", "-c", command_str).Output()
    if err!=nil{
        fmt.Println(err)
        return
    }
    fmt.Println(o)

   
    // t := exec.Command("echo", strconv.Itoa(i), ">", "this.pid")
    // result, _:= t.Output()
    // fmt.Println(string(result) )
}