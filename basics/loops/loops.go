// Learned that '\n' and `\n` does not work.
package loops

import (
	"fmt"
)

func Run() {
	// For
	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", i)
	}
	fmt.Print("\n")
	// For in
	list := [...]rune{'a', 'b', 'c', 'd', 'e', 'f'}
	for idx, elem := range list {
		fmt.Printf("(Idx: %d Elem: '%c'), ", idx, elem)
	}
	fmt.Print("\x0a")
	// While loops do not exist, instead use for with 1 arg
	i := 0
	for i < 10 {
		fmt.Printf("%d, ", i)
		i++
	}
}
