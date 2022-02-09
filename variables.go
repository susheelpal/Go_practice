package main

import (
	"fmt"
)

var a = 10
var b, c, d = 10, 20, "D"

var var2 int32 = 32

func main() {
	var1 := "variable 1"
	fmt.Println("a =", a)
	fmt.Println("b, c and d are", b, c, d)
	fmt.Printf("b = %d,c = %d and d = %s\n", b, c, d)
	fmt.Println("var1 =", var1)
	fmt.Println("var2 =", var2)
	fmt.Printf("a= %T, b= %T, c= %T, d= %T, var2= %T", a, b, c, d, var2)
}
