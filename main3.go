package main

import (
    "fmt"
    "math"
)

func calculateY(a, b, x float64) float64 {
    numerator := 1 + math.Pow(math.Log10(x/a), 2)
    denominator := b - math.Exp(x/a)
    if math.Abs(denominator) < 1e-10 { // проверка деления на почти нулевое число
        return math.Inf(1)
    }
    return numerator / denominator
}

// Функция для вывода результата в удобочитаемом виде
func printResults(results []float64) {
    for i, val := range results {
        fmt.Printf("Result %d: %.4f\n", i+1, val)
    }
}

func main() {
    a := 2.0
    b := 0.95
    xn := 1.25
    xk := 2.75
    dx := 0.3

    // Создание массива значений x
    var xValuesA []float64
    currentX := xn
    for currentX <= xk {
        xValuesA = append(xValuesA, currentX)
        currentX += dx
    }

    // Вычисление значений функции
    resultsA := make([]float64, len(xValuesA))
    for i, x := range xValuesA {
        resultsA[i] = calculateY(a, b, x)
    }

    fmt.Println("Task A Results:")
    printResults(resultsA)

    // Переходим к задаче B
    xValuesB := []float64{2.2, 3.78, 4.51, 6.58, 1.2}
    resultsB := make([]float64, len(xValuesB))
    for i, x := range xValuesB {
        resultsB[i] = calculateY(a, b, x)
    }

    fmt.Println("\nTask B Results:")
    printResults(resultsB)
}