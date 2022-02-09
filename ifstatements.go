// decision statements
package main

import (
	"fmt"
)

/* func main() {
	a := 10
	b := 20

	if a > b {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("a is not greater than b")
	}
}
*/

func main() {
	if 10 == 10 {
		if 20 == 20 {
			if 30 == 40 {
				fmt.Println("yes!! 30 is equal to 40")
			} else {
				fmt.Println("nope!! 30 is not equal of 40")
			}
		} else {
			fmt.Println("both values are not equal")
		}
	} else {
		fmt.Println("both values are not equal")
	}
}
