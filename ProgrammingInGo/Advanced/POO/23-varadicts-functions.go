package main

import (
    "fmt"
)

func add(values ...int) int {
    total := 0
    for _, value := range values {
        total += value
    }
    return total
}

func printNames (names ...string) {
    for _, value := range names {
        fmt.Println(value)
    }
}

func getValues(num int) (double int, triple int, quad int) {
    double = 2 * num
    triple = 3 * num
    quad = 4 * num
    return
}

func main() {
    fmt.Println(add(1))
    fmt.Println(add(1, 2))
    fmt.Println(add(1, 2, 3))
    fmt.Println(add(1, 2, 3, 4))
    fmt.Println(add(1, 2, 3, 4, 5))
    fmt.Println(add(1, 2, 3, 4, 5, 6))
    fmt.Println(add(1, 2, 3, 4, 5, 6, 7))

    printNames("Jhon", "Pedro", "Davis", "Shakira")
    fmt.Println(getValues(2))
}
