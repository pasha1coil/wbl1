package main

import (
	"errors"
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5} // определяем массив
	i := 4                      // определяем i
	newarr, err := remove(arr, i)
	if err != nil {
		fmt.Println(err) // выводим ошибку
	}
	fmt.Printf("new arr:%d\n", newarr) // выводим новый массив

	fmt.Println(append(arr[:i], arr[i+1:]...)) // если порядок важен, то так
}

// функция удаления i элемента из массива
func remove(arr []int, i int) ([]int, error) {
	if i > len(arr)-1 { // проверяем на корректность i элемент
		return nil, errors.New("I > len(arr)-1")
	}
	arr[i] = arr[len(arr)-1]     // меняем i элемент и последний местами
	return arr[:len(arr)-1], nil // возвращаем все элементы массива кроме последнего
}
