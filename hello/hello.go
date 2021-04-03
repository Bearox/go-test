package hello
import "fmt"
import "rsc.io/quote"
import "github.com/Bearox/go-test/greetings"
func main() {
	fmt.Println("hello world.")
	fmt.Println(quote.Go())
    fmt.Println(greetings.Hello("weiwen"))
}
