package conditionals

import (
	"fmt"
)

func Run() {
	if 1 > 2 {
		fmt.Print("1 > 2")
	} else {
		fmt.Print("1 <= 2")
	}
}
