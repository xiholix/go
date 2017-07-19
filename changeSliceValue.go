package main
import(
    "fmt"
)

func changeValue(a *[] int){
    *a= []int{3,4,5,6}
}

func testAppend(){
    a := make([]int, 3,3)
    a[0] = 1
    a[1] = 2
    a[2] = 3
    b := a
    c := &a
    *c = append(*c, 5)
    *c = append(*c, 9)
    *c = append(*c, 9)
    *c = append(*c, 9)
    fmt.Println(b)
    fmt.Println(*c)
}

func main(){
    a := []int{1,2,3}  
    t := &a
    a = []int{3,4,5,6}   //对slice赋值会改变它的指针，长度等属性，之前由它赋值的得到的slice内容却不会改变，因为它对应的属性不会改变
    // changeValue(&a)
    fmt.Println(t)
    fmt.Println(a)

    testAppend()
    
}