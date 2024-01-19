package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int) []int {
	//Если в массиве 1 или меньше элементов
	if len(arr) < 2 {
		return arr
	}
	//Создается правая и левая часть массива
	left, right := 0, len(arr)-1
	//Выбираем опорный элемент (рандомно)
	p := rand.Int() % len(arr)
	//Перемещаем опорный элемент в конец массива
	arr[p], arr[right] = arr[right], arr[p]
	//Разбиваем массив
	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	//Перемещаем опорный элемент на его конечное место
	arr[left], arr[right] = arr[right], arr[left]
	//Рекурсивно применяем сортировку к подмассивам
	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func task16() {
	arr := []int{10, 23, 1, 4, 6, 130, 24, 45, 50}

	fmt.Println("Начальный массив:", arr)
	fmt.Println("Отсортированный массив:", quickSort(arr))
}
