package main

import (
	"fmt"
	"sync"
)

func main() {
	var arr = []int{2, 4, 6, 8, 10}
	ch := make(chan int)    //канал для приема квадратов чисел
	var wg sync.WaitGroup   // вэйт группа для определения количества воркеров
	var resultMu sync.Mutex // мьютекс чтобы избежать гонки за данными
	wg.Add(len(arr))
	for _, i := range arr {
		go Qua2(i, &resultMu, &wg, ch) // передаем адреса переменных
	}
	go func() { // ждем выполнения в отдельной горутине
		wg.Wait()
		close(ch) // закрываем канал по выполнению всех воркеров
	}()
	sum := 0
	for i := range ch { //цикл для сложения
		sum += i
	}
	fmt.Println(sum)

}

// функция записи в кнал данных о квадрате числа
func Qua2(val int, mu *sync.Mutex, wg *sync.WaitGroup, ch chan int) {
	mu.Lock()       //лочим
	ch <- val * val //записываем в канал
	mu.Unlock()     //анлочим
	wg.Done()       //уведомляем о завершении воркера
}
