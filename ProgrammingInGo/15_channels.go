package main

import "fmt"

func say(msg string, canal chan string) {
    canal <- msg
}

func main() {
    canal := make(chan string, 1)

    fmt.Println("Hello")

    go say("Good bye", canal)

    fmt.Println(<-canal)
}
