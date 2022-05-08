package main

import "fmt"

type Person struct {
    name string
    age int
}

type Employee struct {
    id int
}

type FullTimeEmploye struct {
    Person
    Employee
}

func GetMessage(person Person) {
    fmt.Printf("with a age: %d and name %s\n", person.age, person.name)
}

func main() {
    ftEmployee := FullTimeEmploye {}
    ftEmployee.name = "Jhon"
    ftEmployee.age = 29
    ftEmployee.id = 5
    fmt.Printf("%v\n",ftEmployee)
    // Error
    // GetMessage(ftEmployee)
}
