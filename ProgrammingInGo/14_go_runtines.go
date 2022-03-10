package main

import (
    "fmt"
    //"time"
    "sync"
)

func sayHello(msg string, wg *sync.WaitGroup){
    defer wg.Done()
    fmt.Println(msg)
}

func main() {
    var wg sync.WaitGroup
    
    fmt.Println("Hello")

    wg.Add(2)
    go sayHello("Word", &wg)
        
    go func(wgp *sync.WaitGroup) {
        defer wgp.Done()
        fmt.Println("Good Bye")
    }(&wg)

    wg.Wait()

    // No recomend
    // time.Sleep(time.Second * 1)
}
