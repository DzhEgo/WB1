package main

import (
	"fmt"
	"sync"
)

// Создаем функцию записи в мапу
func WriteMap(id int, mutex *sync.Mutex, data map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	//Генерируем ключ для горутин
	key := fmt.Sprintf("Ключ%d", id)
	//Захватываем мутекс для конкуретной записи
	mutex.Lock()
	defer mutex.Unlock()
	//Записываем данные
	data[key] = id
	fmt.Printf("Горутина %d записала в мапу: %s -> %d\n", id, key, id)
}

func task7() {
	//Создаем мапу
	data := make(map[string]int)
	//Создаем мутекс для доступа в мапу
	var mutex sync.Mutex
	//Создаем WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup
	//Кол-во горутин
	num := 3
	//Счетчик
	wg.Add(num)
	//Запускаем горутину
	for i := 0; i < num; i++ {
		go WriteMap(i, &mutex, data, &wg)
	}
	//Ожиданаем завершения
	wg.Wait()

	fmt.Println("Данные мапы: ", data)
}
