package main

import "fmt"

type Employee struct {
    id int
    name string
    vacation bool
}

func (employee *Employee) SetId(id int) {
    employee.id= id;
}

func (employee *Employee) SetName(name string) {
    employee.name = name
}

func (employee *Employee) SetVacation(vacation bool) {
    employee.vacation = vacation
}

func (employee *Employee) GetName() string {
    return employee.name
}

func (employee *Employee) GetId() int {
    return employee.id
}

func (employee *Employee) GetVacation() bool {
    return employee.vacation
}

func NewEmployee(id int, name string, vacation bool) *Employee {
    return &Employee {
        id: id,
        name: name,
        vacation: vacation,
    }
}

func main() {
    employee := Employee {
        id: 1,
        name: "Davis",
        vacation: true,
    }
    fmt.Printf("%v\n", employee)
    employee.SetName("Jhon")
    employee.SetId(20)
    employee.vacation = false

    employee2 := new(Employee)
    employee3 := NewEmployee(2, "Michael", false)

    fmt.Printf("\n");
    fmt.Println(employee.GetId())
    fmt.Println(employee.GetName())
    fmt.Println(employee.GetVacation())
    fmt.Printf("%v\n", employee);
    fmt.Printf("%v\n", *employee2);
    fmt.Printf("%v\n", *employee3);
}
