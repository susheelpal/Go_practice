package main

import (
	"fmt"
)

// OPERATORS
/*
arithmatic OPERATORS ----------- +,-,*,/,%
logical OPERATORS -------------- and,or,not
bitwise OPERATORS -------------- &,|,^,<<,>>,^&
relatioal OPERATORS ------------ >,<,<=,>=,!=,==
assignment OPERATORS ----------- =,+=,-=,*=,/=,&=,|=,^=
misc OPERATORS ----------------- <-, *, &
*/

func main() {
	// arithmatic OPERATORS ----------- +,-,*,/,%
	fmt.Println("1 + 2 =", 1+2)
	// logical OPERATORS -------------- and,or,not
	fmt.Println("true and false =", true && false)
	fmt.Println("not true =", !true)
	// bitwise OPERATORS -------------- &,|,^,<<,>>,^&
	fmt.Println("3 bitwise and 4 =", 3&4)
	fmt.Println("3 bitwise or 4 =", 3|4)
	// relatioal OPERATORS ------------ >,<,<=,>=,!=,==
	fmt.Println("10 > 20", 10 > 20)
	fmt.Println("10 >= 20", 10 >= 20)
	fmt.Println("10 > 10", 10 > 10)
	fmt.Println("10 >= 10", 10 >= 10)
	// assignment OPERATORS ----------- =,+=,-=,*=,/=,&=,|=,^=
	// var a int
	// a = 10
	a := 10
	fmt.Println("a =", a)
	a += 10
	fmt.Println("a =", a)
	// misc OPERATORS ----------------- <-, *, &
	b := 20
	p := &b
	fmt.Println("accessing via pointer", *p)
}
