package main

import (
    "time"
)

type Person struct {
    DNI string
    Name string
    Age int
}

type Employee struct {
    Id int
    Position string
}

type FullTiempEmployee struct {
    Employee
    Person
}

var GetPersonByDni = func(dni string) (Person, error) {
    time.Sleep(5 *time.Second)
    return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
    time.Sleep(5 *time.Second)
    return Employee{}, nil
}

func GetFullTiempEmployee(id int, dni string) (FullTiempEmployee, error) {
    var ftEmployee FullTiempEmployee

    e, err := GetEmployeeById(id)
    if err != nil {
        return ftEmployee, err
    }
    ftEmployee.Employee = e

    p, err := GetPersonByDni(dni)
    if err != nil {
        return ftEmployee, err
    }
    ftEmployee.Person = p
    
    return ftEmployee, nil
}
