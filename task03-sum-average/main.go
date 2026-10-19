package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите число N: ")
	fmt.Scan(&n)

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	var avg float64
	if n > 0 {
		avg = float64(sum) / float64(n)
	} else {
		avg = 0
	}

	fmt.Printf("Сумма чисел от 1 до %d: %d\n", n, sum)
	fmt.Printf("Среднее значение: %.2f\n", avg)
}