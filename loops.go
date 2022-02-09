// Loops in golang
package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i + 1)
	}
	fmt.Println("-------while loop-------")
	j := 1
	for j < 5 {
		fmt.Println(j)
		j += 1
	}

	fmt.Println("-----------for using range-----------")

	for m, n := range []int{1, 2, 3, 4, 5} {
		fmt.Println(m, n)
	}

	fmt.Println("-----------for using range(do not print key)-----------")
	for _, n := range []int{1, 2, 3, 4, 5} {
		fmt.Println(n)
	}
	fmt.Println("-----------infinite loop-----------")
	/* for {
		fmt.Println("-----------infinite loop-----------")
	} */

}
