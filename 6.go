package main

import (
	"sync"
	"time"
)

// С помощью WaitGroup
func task6() {
	//Используем, для ожидания завершения горутины
	var wg sync.WaitGroup
	//Счетчик
	wg.Add(1)
	//Запускаем горутину
	go func(wg *sync.WaitGroup) {
		//Уменьшаем счетчик при завершении горутины
		defer wg.Done()
		//...
	}(&wg)
	//Ожидаем завершения горутин
	wg.Wait()
}

// С помощью сигнала
func task6_1() {
	//Создаем канала для передачи сигнала
	stopChan := make(chan struct{})
	//Запускаем горутину
	go func(stopChan <-chan struct{}) {
		for {
			select {
			case <-stopChan:
				//Получаем сигнал завершения и выходим из горутины
				return
			default:
				//Горутина работает
			}
		}
	}(stopChan)
	//Отправляем сигнал о завершении
	close(stopChan)
}

// С помощью таймера
func task6_3() {
	//Создаем таймер на 5 секунд
	timer := time.NewTimer(5 * time.Second)

	//Запускаем горутину
	go func() {
		for {
			select {
			case <-timer.C:
				//Истекает тацмер - выходим из горутины
				return
			default:
				//Работа горутины
			}
		}
	}()
}
