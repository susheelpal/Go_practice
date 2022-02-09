package main

import (
	"fmt"
)

func main() {
	value := func() {
		fmt.Print("anonymous function")
	}

	value()
}
