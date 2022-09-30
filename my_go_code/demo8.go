package main
import (
    "fmt"
    "time"
)
var(
    mymap=make(map[int]int ,10)
)

func test(n int){
    res:=1
    for i:=1;i<=n;i++{
        res+=i;


    }
    mymap[n]=res
}
func main(){
    for i:=1;i<=1000;i++{
        go test(i)
    }
	time.Sleep(1 * time.Second)
    for i,v:=range mymap{
        fmt.Printf("map[%d]=%d\n",i,v)
    }
}