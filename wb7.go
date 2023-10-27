package main

import (
	"fmt"
	"sync"
)

// структура для мапы, сама мапа и мьютекс
type Mapa struct {
	mu sync.Mutex
	mp map[int]int
}

// метод стуктуры, используется мьютекс для того чтобы избежать гонки данных
func (s *Mapa) Add(key int, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.mp[key] = value
}

func main() {
	// инициализируем мапу
	Mapa := Mapa{
		mp: make(map[int]int),
	}
	// вэйтгрупп для группы горутин
	var wg sync.WaitGroup
	numOfGoroutines := 10
	wg.Add(numOfGoroutines)

	//
	for i := 0; i < numOfGoroutines; i++ {
		go func(i int) {
			Mapa.Add(i, i) // запускаем метод
			wg.Done()      // завершаем горутину
		}(i)
	}

	wg.Wait() //ждем выполнения всех воркеров

	fmt.Println(Mapa.mp)
}
