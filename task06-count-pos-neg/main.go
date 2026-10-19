package main

import "fmt"

func main() {
	var n int
	fmt.Print("Сколько чисел вы хотите ввести? ")
	fmt.Scan(&n)

	positive := 0
	negative := 0
	zeros := 0

	for i := 0; i < n; i++ {
		var num int
		fmt.Printf("Введите число %d: ", i+1)
		fmt.Scan(&num)

		if num > 0 {
			positive++
		} else if num < 0 {
			negative++
		} else {
			zeros++
		}
	}

	fmt.Printf("Положительных чисел: %d\n", positive)
	fmt.Printf("Отрицательных чисел: %d\n", negative)
	fmt.Printf("Нулей: %d\n", zeros)
}