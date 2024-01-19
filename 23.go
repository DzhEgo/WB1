package main

import "fmt"

func task23() {
	arr := []int{1, 2, 3, 4}
	i := 2
	//Удаляем i элемент
	arr = append(arr[:i], arr[i+1:]...)

	fmt.Println(arr)
}
