package main
import (
    "fmt"
    "time"
)
func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s, (i+1)*100)
    }
}
func say2(s string ,ch1 chan int) {
    for i := 0; i < 5; i++ {
        time.Sleep(150 * time.Millisecond)
        fmt.Println(s, (i+1)*150)
    }
	ch1<-0
	close(ch1)
}
func main() {
	ch :=make(chan int)
    go say2("world",ch)
    say("hello")
	fmt.Println(<-ch)

}