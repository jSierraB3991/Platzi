package main

import "fmt"

func say(msg string, canal chan<- string) {
    canal <- msg
}

func main() {
    canal := make(chan string,2)

    // fmt.Println("Hello")
    canal <- "Message 1"
    canal <- "Message 2"
    // go say("Good bye", canal)

    fmt.Println(len(canal), cap(canal))
    //fmt.Println(<-canal)
    close(canal)

    for msg := range canal {
        fmt.Println(msg)
    }

    // Error, this channel is a close
    // canal <- "Message 3"

    /// SELECT
    email := make(chan string)
    email2 := make(chan string)
    
    go say("message1", email)
    go say("message2", email2)

    for i := 0; i < 2; i++ { 
        select {
        case m1 := <- email:
            fmt.Println("mail recive of email1:", m1)
        case m2 := <- email2:
            fmt.Println("mail recive of email2: ", m2)
        }
    }
}
