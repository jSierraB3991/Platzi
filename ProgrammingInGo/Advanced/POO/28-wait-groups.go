package main

import (
    "fmt"
    "sync"
    "time"
)

func doSmething(id int, wg *sync.WaitGroup) {
    defer wg.Done()

    fmt.Printf("Sleeping %d\n", id)
    time.Sleep(2 * time.Second)
    fmt.Printf("Wake Up %d\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i:=0; i<10; i++ {
        wg.Add(1)
        go doSmething(i, &wg)
    }
    wg.Wait()
}
