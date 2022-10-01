package main

import (
    "fmt"
)

type IProduct interface {
    SetStock(stock int)
    GetStock() int
    SetName(name string)
    GetName() string
}

type Computer struct {
    stock int
    name string
}

func (c *Computer) GetStock() int {
    return c.stock
}
func (c *Computer) SetStock(stock int) {
    c.stock = stock
}

func (c *Computer) GetName() string {
    return c.name
}
func (c *Computer) SetName(name string) {
    c.name = name
}

type Laptop struct {
    Computer
}

func NewLaptop() IProduct {
    return &Laptop {
        Computer: Computer{
            name: "Laptop computer",
            stock: 25,
        },
    }
}

type Desktop struct {
    Computer
}

func NewDesktop() IProduct  {
    return &Desktop {
        Computer: Computer {
            name: "Desktop Computer",
            stock: 35,
        },
    }
}

func GetComputerFactory(typePC string) (IProduct, error){
    if typePC == "laptop" {
        return NewLaptop(), nil
    }
    if typePC == "desktop" {
        return NewDesktop(), nil
    }

    return nil, fmt.Errorf("Invalid Computer Type")
}

func printNameAndStock(p IProduct){
    fmt.Printf("product name: %s, with stock: %d \n", p.GetName(), p.GetStock())
}

func main(){
    desktop, err := GetComputerFactory("desktop")
    if err == nil {
        printNameAndStock(desktop)
    }

    laptop, err := GetComputerFactory("laptop")
    if err == nil {
        printNameAndStock(laptop)
    }
}
