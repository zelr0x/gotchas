package main

import (
	"fmt"
)

/*
--- kinda ok ---
4
[2 3 4 5]
4
8
x = 100,	y = 100,	z = 5000

--- wtf ---
4
[2 3]
4
4
x = 5000,	y = 5000,	z = 5000
--- wtf2 ---
5
[2 3 4 5]
5
5
x = 5000,	y = 5000,	z = 5000
*/
func main() {
	fmt.Println("--- kinda ok ---")
	kindaOk()
	fmt.Println()
	fmt.Println("--- wtf ---")
	wtf()
	fmt.Println("--- wtf2 ---")
	wtf2()
}

func kindaOk() {
	test(1, 5)
}

func wtf() {
	test(1, 3)
}

func wtf2() {
	x := [...]int{1, 2, 3, 4, 5, 6}
	y := x[1:5]
	fmt.Println(cap(y))
	fmt.Println(y)
	y[0] = 100
	fmt.Println(cap(y))
	z := append(y, 1000)
	fmt.Println(cap(z))
	z[0] = 5000
	fmt.Printf("x = %v,\ty = %v,\tz = %v\n", x[1], y[0], z[0])
}

func test(begin, end uintptr) {
	x := [...]int{1, 2, 3, 4, 5}
	y := x[begin:end]
	fmt.Println(cap(y))
	fmt.Println(y)
	y[0] = 100
	fmt.Println(cap(y))
	z := append(y, 1000)
	fmt.Println(cap(z))
	z[0] = 5000
	fmt.Printf("x = %v,\ty = %v,\tz = %v\n", x[1], y[0], z[0])
}
