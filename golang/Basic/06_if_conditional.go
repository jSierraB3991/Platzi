package main

import "fmt"

func main() {
    value1 := 1
    value2 := 2
   
    fmt.Println("if and else")
    if value1 == 1 {
        fmt.Println("The value is 1")
    } else {
        fmt.Println("The value is not 1")
    }

    fmt.Printf("\nFloodgates logic if\n")
    if value1 == 4 && value2 == 2 {
        fmt.Println("The value1 is 4 and value2 is 2")
    }

    fmt.Printf("\nFloodgates logic or\n")
    if value1 == 3 || value2 == 2 {
        fmt.Println("The value1 is 3 or value2 is 2")
    }
}
