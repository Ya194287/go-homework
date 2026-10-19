package main

import "fmt"

// Функция проверяет, является ли число простым
func isPrime(n int) bool {
	if n <= 1 {
		return false // 1, 0 и отрицательные числа не являются простыми
	}
	// Достаточно проверить делители до корня из n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false // Нашли делитель, значит число не простое
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("Введите число для проверки: ")
	fmt.Scan(&n)

	if isPrime(n) {
		fmt.Printf("%d — простое число\n", n)
	} else {
		fmt.Printf("%d — не простое число\n", n)
	}
}