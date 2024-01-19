package main

import (
	"fmt"
	"sync"
)

// То же самое, что и в предыдущем задании, за исключением одной вещи. Смотрим ниже...
func squarePlus(wg *sync.WaitGroup, numbers <-chan int, squares chan<- int) {
	for number := range numbers {
		squares <- number * number
	}
	wg.Done()
}

func task3() {
	numbers := []int{2, 4, 6, 8, 10}

	numbersChan := make(chan int, len(numbers))
	squaresChan := make(chan int, len(numbers))

	var wg sync.WaitGroup

	for i := 0; i < len(numbers); i++ {
		wg.Add(1)
		go squarePlus(&wg, numbersChan, squaresChan)
	}

	for _, number := range numbers {
		numbersChan <- number
	}
	close(numbersChan)
	wg.Wait()
	close(squaresChan)

	//Создаем переменную, которая будет отвечать за сумму всех квадратов
	sum := 0
	//Пробегаемся по всем числам и суммируем их в переменную sum
	for square := range squaresChan {
		sum += square
	}
	//Выводим результат
	fmt.Println("Сумма:", sum)
}
