package main
import(
    "fmt"
    "sync"
    "time"
)

var a = []int{1,2,3,4}
var count = 5

func addOneElement(rw *sync.RWMutex, a *[]int){
     for{
         time.Sleep(3*time.Second)
         rw.Lock()
         time.Sleep(3*time.Second)
         *a = append(*a, count)
         count += 1
         rw.Unlock()

     }
}


func printElem(rw *sync.RWMutex, i int, a *[]int){  // 必须是指针，否则无法锁定， a也必须是指针，否则append的值无法加到原来的slice上。
    for{
        rw.RLock()
        time.Sleep(time.Second)
        fmt.Println("%d is %v\n",i , a)
        rw.RUnlock()
    }
}

func main(){
    rw := sync.RWMutex{}
    go addOneElement(&rw, &a)
    for i:=0; i<2; i++{
        go printElem(&rw, i, &a)
    }
    for{
        time.Sleep(time.Second*5)
    }
}