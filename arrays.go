package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("type = %T", arr1)
	fmt.Printf("\ntype = %T", arr2)
}
