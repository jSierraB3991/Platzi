package main

import "fmt"

func main() {

    fmt.Println("Tradicional for: ")
    for i := 0; i <= 10; i++ {
        fmt.Printf("%d", i)
    }

    fmt.Printf("\n\n For While: \n")
    counter := 10
    for counter >= 0 {
        fmt.Printf("%d", counter)
        counter--
    }

    fmt.Println("\n\nFor Forever")
    //counterForever := 0
    //for {
        //fmt.Printf("%", counterForever)
        //counterForever++
    //}
    fmt.Println("End Program")
}
