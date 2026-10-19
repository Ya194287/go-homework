package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите число N: ")
	fmt.Scan(&n)

	// Положительное / отрицательное
	if n > 0 {
		fmt.Println("Число положительное")
	} else if n < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Число равно нулю")
	}

	// Чётное / нечётное
	if n%2 == 0 {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число нечётное")
	}

	// Делится на 3 и на 5
	if n%3 == 0 && n%5 == 0 {
		fmt.Println("Число делится и на 3, и на 5")
	} else if n%3 == 0 {
		fmt.Println("Число делится на 3, но не на 5")
	} else if n%5 == 0 {
		fmt.Println("Число делится на 5, но не на 3")
	} else {
		fmt.Println("Число не делится ни на 3, ни на 5")
	}
}