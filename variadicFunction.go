package main

import (
	"fmt"
)

func add(ele ...int) int {
	sum := 0
	for _, element := range ele {
		sum += element
	}
	return sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := add(arr...)
	fmt.Println("sum =", res)
}
