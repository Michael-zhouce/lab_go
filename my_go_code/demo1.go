package main
import(
	"fmt"
)

func AddUpper(n int ) func (int) int{
	
	return  func(x int) int {
		n=n+x
		return n
	}
}

func main (){
	f:=AddUpper(1)
	fmt.Println(f(4))
}