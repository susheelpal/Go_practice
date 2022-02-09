// switch case statement

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Switch case statement")

	i := 5
	switch i {
	case 1:
		fmt.Println("value is", i)
	case 2:
		fmt.Println("value is", i)
	case 3:
		fmt.Println("value is", i)
	case 4, 5:
		fmt.Println("value is", i)
	default:
		fmt.Println("value is default")
	}
}
