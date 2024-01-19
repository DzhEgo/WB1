package main

import "fmt"

// Синтаксис
func task13() {
	a := 12
	b := 13
	a, b = b, a

	fmt.Println(a)
	fmt.Println(b)
}

// Математика
func task13_1() {
	a := 12
	b := 13

	a += b
	b = a - b
	a = a - b

	fmt.Println(a)
	fmt.Println(b)
}
