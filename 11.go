package main

import "fmt"

// Создаем функцию для проверки на пересечения множеств
func peres(set1, set2 map[int]bool) map[int]bool {
	//Создаем мапу, для добавления значений, которые пересекаются
	res := make(map[int]bool)
	//Проходимся по элементам первого множества
	for e := range set1 {
		if set2[e] {
			//Если во втором множестве значение, как и в первом, то добавляем его в результат
			res[e] = true
		}
	}
	return res
}

func task11() {
	//Создаем множества
	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true}
	//Использвуем функцию
	res := peres(set1, set2)
	//Выводим результат
	fmt.Println("Пересечение множеств:", res)
}
