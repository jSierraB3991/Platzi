package main

import "fmt"

func variousValue(a, b int, c string) int {
    fmt.Printf("message sending es: %s \n", c)
    return a * b
}

func variousReturn(a int) (c, b int) {
    return a, a*3
}

func main() {
    value := variousValue(2, 3, "hola prro")
    fmt.Printf("The value return is %d \n", value)

    value2, value3 := variousReturn(2)
    fmt.Printf("Two values return is %d and %d \n", value2, value3)

    // push _ if no required second value
    value4, _ := variousReturn(10)
    fmt.Printf("Alone value in two returns is %d \n", value4)
}
