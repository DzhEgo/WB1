package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, el int) int {
	//Используем sort.Search для бинарного поиска. Метод принимает длину массива и функцию, которая возвращает true, если
	//Значение в i больше или равно исходному значению el
	ind := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= el
	})
	if ind < len(arr) && arr[ind] == el {
		//Элемент найден
		return ind
	}
	//Элемент не найден
	return -1
}

func task17() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	el := 9

	fmt.Println("Массив:", arr)
	fmt.Printf("Индекс элемента %d: %d\n", el, binarySearch(arr, el))
}
