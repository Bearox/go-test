package main
import (
	"fmt"
	"time"
)
import "github.com/Bearox/go-test/greetings"

func main() {
	fmt.Println("hello world.")
    fmt.Println(greetings.Hello("weiwen"))
	fmt.Println(5 & 1)
	fmt.Println(10 & 1)
	fmt.Println(time.Now().Second() & 1)
	fmt.Println(time.Now().Add( - time.Minute * 10))
}
