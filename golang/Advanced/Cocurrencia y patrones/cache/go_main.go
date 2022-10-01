package main

import (
    "fmt"
    "log"
    "sync"
    "time"
)

func Fibonacci(n int) int {
    if n <= 1{
        return n
    }
    return Fibonacci(n-1) + Fibonacci(n-2)
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
    value interface{}
    err error
}

type Memory struct {
    f Function
    cache map[int]FunctionResult
    mutex sync.Mutex
}

func NewCache(f Function) *Memory {
    return &Memory {
        f: f,
        cache: make(map[int]FunctionResult),
    }
}

func (m *Memory) Get(key int) (interface{}, error) {
    m.mutex.Lock()
    result, exists := m.cache[key]
    m.mutex.Unlock()
    if ! exists {
        m.mutex.Lock()
        result.value, result.err = m.f(key)
        m.cache[key] = result
        m.mutex.Unlock()
    }

    return result.value, result.err

}

func GetFibonacci(n int) (interface{}, error) {
    return Fibonacci(n), nil
}

func main() {
    cache := NewCache(GetFibonacci)
    fibo := []int{42, 40, 41, 42, 38}
    var wg sync.WaitGroup

    for _, fib := range fibo {
        wg.Add(1)
        go func(index int) {
            defer wg.Done()
            start := time.Now()
            value, err := cache.Get(index)
            if err != nil {
                log.Println(err)
            }
            fmt.Printf("Calculate: %d, Time: %s, Result: %d\n", index, time.Since(start), value)
        }(fib)
    }
    wg.Wait()
}
