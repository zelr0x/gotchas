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
*/
func main() {
	fmt.Println("--- kinda ok ---")
	kindaOk()
	fmt.Println()
	fmt.Println("--- wtf ---")
	wtf()
}

// This one behaves as expected in go.
func kindaOk() {
	test(1, 5)
}

// This one shows that capacity of a slice may be greater
// than its length even before any appends!
func wtf() {
	test(1, 3)
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
