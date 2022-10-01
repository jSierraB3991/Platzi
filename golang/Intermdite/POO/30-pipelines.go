package main

import (
    "fmt"
)

// channel of read
func Generator(c chan<- int) {
    for i:=0; i<=10; i++ {
        c <- i
    }
    close(c)
}

// cin => channel of read, cout => cout of writer
func Double(cin <-chan int, cout chan<- int) {
    for value := range cin {
        cout <- 2 * value
        // ERROR
        // cin <- 2 * value
    }
    close(cout)
}

// channel of read
func Print(c <-chan int) {
    for value := range c {
        fmt.Printf("%d ", value)
    }
}


func main() {
    cgenerator := make(chan int)
    cdouble := make(chan int)

    go Generator(cgenerator)
    go Double(cgenerator, cdouble)
    Print(cdouble)
    fmt.Println("")
}
