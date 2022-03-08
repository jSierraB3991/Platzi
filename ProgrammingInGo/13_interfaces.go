package main

import "fmt"

type figure2d interface {
    area() float64
}

type frame struct {
    base float64
}

type rectangle struct {
    base float64
    heigh float64
}

func (c frame) area() float64 {
    return c.base * c.base
}

func (r rectangle) area() float64 {
    return r.base * r.heigh
}

func calculate(figure figure2d) {
    fmt.Println("área: ", figure.area())
}

func main() {
    myFrame := frame { base: 2 }
    fmt.Println("frame: ", myFrame)

    myRectangle := rectangle { base: 2, heigh: 4 }
    fmt.Println("rectangle: ", myRectangle)
    
    calculate(myRectangle)
    calculate(myFrame)
}
