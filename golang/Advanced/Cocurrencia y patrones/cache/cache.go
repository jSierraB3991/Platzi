package main

import (
    "fmt"
    "sync"
    "time"
)


func Fibonacci(n int) int {
    fmt.Printf("Calculate expensive Fibonacci for %d\n", n)
    time.Sleep(5 * time.Second)
    return n
}

type Service struct {
    InProgress map[int]bool
    IsPending map[int] []chan int
    Mutex sync.RWMutex
}

func (s *Service) Work(job int) {
    s.Mutex.RLock()
    exists := s.InProgress[job]
    if exists {
        s.Mutex.RUnlock()
        response := make(chan int)
        
        defer close(response)
        s.Mutex.Lock()
        s.IsPending[job] = append(s.IsPending[job], response)
        s.Mutex.Unlock()
        fmt.Printf("Waiting for response job: %d\n", job)
        res := <-response
        fmt.Printf("response value is recibing: %d\n", res)
        return
    }
    s.Mutex.RUnlock()
    
    s.Mutex.Lock()
    s.InProgress[job] = true
    s.Mutex.Unlock()

    fmt.Printf("calculate fibonnaci for job %d\n", job)
    result := Fibonacci(job)
    s.Mutex.RLock()
    pendingWorkers, exists := s.IsPending[job]
    s.Mutex.RUnlock()

    if exists {
        for _, pendingWorker := range pendingWorkers {
            pendingWorker <- result
        }
        fmt.Printf("The result send for all workers in pending in job: %d\n", job)
    }

    s.Mutex.Lock()
    s.InProgress[job] = false
    s.IsPending[job] = make([]chan int, 0)
    s.Mutex.Unlock()
}

func NewService() *Service {
    return &Service {
        InProgress: make(map[int]bool),
        IsPending: make(map[int] []chan int),
    }
}

func main() {
    service := NewService()
    jobs := []int {3,4,5,5,4,8,8,8}
    var wg sync.WaitGroup
    for _, value := range jobs {
        wg.Add(1)
        go func(job int) {
            defer wg.Done()
            service.Work(job)

        }(value)
    }
    wg.Wait()
}
