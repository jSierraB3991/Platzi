package main

import (
    "fmt"
    "sync"
)

var (
    balance int = 100
)

func deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
    defer wg.Done()
    lock.Lock()
    b := balance
    balance = b + amount
    lock.Unlock()
}

func Balance(lock *sync.RWMutex) int {
    lock.RLock()
    b := balance
    lock.RUnlock()
    return b
}

func main() {
    var wg sync.WaitGroup
    var lock sync.RWMutex
    fmt.Printf("The balance star is %v\n", Balance(&lock))
    for i:=1; i<=5; i++ {
        wg.Add(1)
        go deposit(100*i, &wg, &lock)
    }

    wg.Wait()
    fmt.Printf("The final of balance is %v\n", Balance(&lock))
}
