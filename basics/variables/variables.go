package variables

import (
	"fmt"
)

func Run() {
	var a int = 1
	b := 2
	c, d := 3, 4

	b = c - a
	d /= b

	fmt.Printf("4 / (3 - 1) = %d", d)
}
