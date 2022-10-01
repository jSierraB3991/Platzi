package main

import (
    "testing"
)

func TestGetFullTiempEmployee(t *testing.T) {
    table := [] struct {
        id int
        dni string
        mockFunc func()
        expected FullTiempEmployee
    }{
        {
            id: 1, dni: "1", mockFunc: func() {
                GetEmployeeById = func(id int) (Employee, error) {
                    return Employee { Id: 1, Position: "CEO",}, nil
                }
                GetPersonByDni = func(dni string) (Person, error) {
                    return Person { DNI: "1", Name: "Jhon", Age: 35,}, nil
                }
            }, expected: FullTiempEmployee {
                Person: Person {
                    Age: 35, DNI: "1", Name: "Jhon",
                },
                Employee: Employee {
                    Id: 1, Position: "CEO",
                },
            },
        },
    }

    originalGetEmployeeByid := GetEmployeeById
    originalGetPersonByDni := GetPersonByDni

    for _, test := range table {
        test.mockFunc()
        fte, error := GetFullTiempEmployee(test.id, test.dni)
        if error != nil {
            t.Errorf("Error geting Employee id %d, dni %s", test.id, test.dni)
        }

        if fte.Age != test.expected.Age {
            t.Errorf("get: %d,expected: %d", fte.Age, test.expected.Age)
        }
        if fte.DNI != test.expected.DNI{
            t.Errorf("get: %s, expected: %s", fte.DNI, test.expected.DNI)
        }
        if fte.Name != test.expected.Name{
            t.Errorf("get: %s, expected: %s", fte.Name, test.expected.Name)
        }
        if fte.Id != test.expected.Id{
            t.Errorf("get: %d, expected: %d", fte.Id, test.expected.Id)
        }
        if fte.Position != test.expected.Position {
            t.Errorf("get: %s, expected: %s", fte.Position , test.expected.Position )
        }
    }
    GetEmployeeById = originalGetEmployeeByid
    GetPersonByDni = originalGetPersonByDni
}
