
package main

import "fmt"

func main() {
    platzi := "Platzi"
    cursosCount := 5
    
    message := fmt.Sprintf("%s is education on line", platzi)
    fmt.Println(message)

    fmt.Printf("%s hve more of %d courses\n", platzi, cursosCount)
    fmt.Printf("Platzi is a variable %T\n", platzi)
    fmt.Printf("Cursos Count is a variable %T\n", cursosCount)
}

