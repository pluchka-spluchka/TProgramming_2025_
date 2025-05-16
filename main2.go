package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(calculate(2.0, 0.95, 2.2))
}

func calculate(a, b, x float64) float64 {
    numerator := 1 + math.Pow(math.Log10(x/a), 2)

    denominator := b - math.Exp(x/a)

    result := numerator / denominator

    return result
}