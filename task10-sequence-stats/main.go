package main

import "fmt"

func main() {
	var n int
	fmt.Print("Сколько чисел вы хотите ввести? ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Количество чисел должно быть больше нуля")
		return
	}

	sum := 0.0
	max := 0
	min := 0
	evenCount := 0
	oddCount := 0

	for i := 0; i < n; i++ {
		var num int
		fmt.Printf("Введите число %d: ", i+1)
		fmt.Scan(&num)

		// Добавляем к сумме
		sum += float64(num)

		// Инициализируем max и min первым введенным числом
		if i == 0 {
			max = num
			min = num
		} else {
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		}

		// Считаем четные и нечетные
		if num%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	avg := sum / float64(n)

	fmt.Println("\n--- Статистика последовательности ---")
	fmt.Printf("Сумма: %.2f\n", sum)
	fmt.Printf("Среднее значение: %.2f\n", avg)
	fmt.Printf("Максимальное значение: %d\n", max)
	fmt.Printf("Минимальное значение: %d\n", min)
	fmt.Printf("Количество чётных чисел: %d\n", evenCount)
	fmt.Printf("Количество нечётных чисел: %d\n", oddCount)
}