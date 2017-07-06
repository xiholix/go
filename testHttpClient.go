package main 

import "fmt"
import "net/http"
import "io"
import "os"
func main(){

	response, err := http.Get("http://www.neu.edu.cn")
	if err==nil{
		fmt.Println(response)
	}else{
		fmt.Println(err)
	}
	
	defer response.Body.Close()
	io.Copy(os.Stdout, response.Body)

}