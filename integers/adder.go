package integers

import (
	"fmt"
)

// Add takes two integers and returns the sum of them.
func Add(x, y int) (result int) {
	result = x + y
	return
}

func main() {
	fmt.Println(Add(2, 2))
}
