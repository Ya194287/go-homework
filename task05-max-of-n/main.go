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

    var num int
    fmt.Print("Введите число 1: ")
    fmt.Scan(&num)

    max := num
    min := num

    for i := 2; i <= n; i++ {
        fmt.Printf("Введите число %d: ", i)
        fmt.Scan(&num)
        if num > max {
            max = num
        }
        if num < min {
            min = num
        }
    }

    fmt.Printf("Максимальное значение: %d\n", max)
    fmt.Printf("Минимальное значение: %d\n", min)
}