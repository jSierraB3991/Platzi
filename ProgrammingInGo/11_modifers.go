package main

import (
    pk "my_package"
    "fmt"
)

func main() {
    var myCar pk.CarPublic
    myCar.Brand = "Ferrari"
    fmt.Println(myCar)

    // Error
    // var myCar2 pk.carPrivate

    pk.PrintMessage("Hello world!!")
    
    // Error
    // pk.printMessage()
}
