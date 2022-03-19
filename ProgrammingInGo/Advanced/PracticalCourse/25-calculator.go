package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

type calc struct {}

func (calc) operate(operation string, operator string) int {
    
    values := strings.Split(operation, operator)
    operator1 := parser_to_number(values[0]) 
    operator2 := parser_to_number(values[1])
    var result int

    switch operator {
        case "+": 
            result = operator1 + operator2
        case "-": 
            result = operator1 - operator2
        case "*": 
            result = operator1 * operator2
        case "/" :
            result = operator1 / operator2
        default:
            fmt.Printf("%s is not support\n", operator)
    }
    return result
}

func parser_to_number (number_string string) int {
    number, _ := strconv.Atoi(number_string)
    return number
}

func read_term() string {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    return scanner.Text()
}

func main() {
    fmt.Print("What is the runction: ")
    operation := read_term()
    fmt.Print("What is the operator: ")
    operator := read_term()
    fmt.Printf("operation: %s, operator: %s \n", operation, operator)
    calc := calc{}
    fmt.Printf("The result of operation is: %d\n", calc.operate(operation, operator))
}
