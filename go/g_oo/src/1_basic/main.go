package main
// https://flaviocopes.com/golang-is-go-object-oriented/
// https://www.ardanlabs.com/blog/2013/07/object-oriented-programming-in-go.html
import (
	"fmt"
)

func main() {
	fmt.Println("Hello OO\n===================")

	cat1 := Cat{"miao1"}
	cat1.Move(2)

	dog := Dog{"wo"}
	dog.Move(3)

	cat2 := Cat{"mi2"}
	cat2.Move(4)

	// cat2.hello()  // cannot be called
}