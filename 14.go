package main

import (
	"fmt"
	"reflect"
)

// Через пакет reflect и метода TypeOf
func task14() {
	var i int
	fmt.Println(reflect.TypeOf(i))

	var str string
	fmt.Println(reflect.TypeOf(str))

	var boo bool
	fmt.Println(reflect.TypeOf(boo))

	inter := make(chan interface{})
	fmt.Println(reflect.TypeOf(inter))
}

// через type в interface
func task14_1(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan interface{}:
		fmt.Println("chan interface{}")
	}
}
