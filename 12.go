package main

import "fmt"

func task12() {
	//Создаем последовательность строк
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	//Создаем мапу, для выявления множеств
	uniq := make(map[string]bool)
	//Проходимся по строкам и выявляем уникальные значения
	for _, el := range arr {
		uniq[el] = true
	}
	//Выводим уникальные значения
	for el := range uniq {
		fmt.Println(el)
	}
}
