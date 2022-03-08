package my_package

// The fist letter to lower case is private
// The fist letter to upper case is public

import "fmt"

// CarPublic is a struct car public
type CarPublic struct {
    Brand string
    Year int
}

type carPrivate struct {
    brand string
    year int
}

// PrintMessage print a message
func PrintMessage(message string) {
    fmt.Println(message)
}

func printMessage(){
    fmt.Println("Hello!!")
}
