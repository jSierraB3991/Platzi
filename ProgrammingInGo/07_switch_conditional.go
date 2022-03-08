package main

import "fmt"

func main() {
    module := 5%2

    switch module {
    case 0:
        fmt.Println("The module is par")
    default:
        fmt.Println("The module is impar")
    }
    
    switch module2 :=  4%2; module2 {
    case 0:
        fmt.Println("The module2 is par")
    default:
        fmt.Println("The module2 is impar")
    }

    value := 20
    switch {
    case value > 100:
        fmt.Println("The value is more to 100")
    case value < 100:
        fmt.Println("The value is low to 100")
    default:
        fmt.Println("The value is 100")
    }
}
