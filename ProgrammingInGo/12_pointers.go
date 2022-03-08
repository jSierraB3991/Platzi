package main

import (
    "fmt"
    mpk "my_package"
)

func main() {
    a := 50
    b := &a

    fmt.Println(a, ", ", b)
    fmt.Println(*b)
    
    fmt.Printf("\nChange value\n")
    *b = 100
    fmt.Println(a, ", ", b)
    fmt.Println(*b)

    fmt.Printf("\nOther example\n")
    myPc := mpk.Pc { Ram: 16, Disk: 20, Brand: "MSI"}
    myPc.Ping()
    fmt.Println(myPc)
    myPc.DuplicateRam()
    fmt.Println(myPc)
    myPc.DuplicateRam()
    fmt.Println(myPc)
}
