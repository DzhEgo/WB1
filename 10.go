package main

import (
	"fmt"
	"math"
	"sort"
)

func group(temperatures []float64, step float64) map[string][]float64 {
	groups := make(map[string][]float64)
	// Сортируем температуры
	sort.Float64s(temperatures)

	for _, temp := range temperatures {
		// Находим ключ (диапазон температур) для текущей температуры
		key := fmt.Sprintf("%.0f-%.0f", math.Floor(temp/step)*step, math.Ceil(temp/step)*step)
		// Добавляем температуру в соответствующую группу
		groups[key] = append(groups[key], temp)
	}
	return groups
}

func task10() {
	//Обозначаем температуры
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	//Обозначаем шаг
	step := 10.0
	//Группируем
	groups := group(temp, step)
	//Выводим
	for key, val := range groups {
		fmt.Printf("Группа %s: %v\n", key, val)
	}
}
