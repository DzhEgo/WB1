package main

import (
	"fmt"
	"sync"
)

// Создаем функцию нахождения квадрата числа и отправки результат в канал. Из numbers достаем числа и записываем в squares
func square(wg *sync.WaitGroup, numbers <-chan int, squares chan<- int) {
	for number := range numbers {
		squares <- number * number
	}
	//После обработки всех чисел с помощью wg.Done() завершаем работу функцию
	wg.Done()
}

func task2() {
	//Инициализируем массив
	numbers := []int{2, 4, 6, 8, 10}
	//Создаем каналы
	numbersChan := make(chan int, len(numbers))
	squaresChan := make(chan int, len(numbers))
	//С помощью WaitGroup ожидаем завершения всех горутин
	var wg sync.WaitGroup
	//Запускаем горутины
	for i := 0; i < len(numbers); i++ {
		//Добавляем горутину в счетчик
		wg.Add(1)
		go square(&wg, numbersChan, squaresChan)
	}
	//Отправляем числа в канал
	for _, number := range numbers {
		numbersChan <- number
	}
	close(numbersChan)
	//Ожидаем завершения всех горутин
	wg.Wait()
	close(squaresChan)
	//Выводим результат
	for square := range squaresChan {
		fmt.Println(square)
	}
}
