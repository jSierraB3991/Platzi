package main

import (
    "fmt"
)

func Fibonnaci(num int) int {
    if num <= 1 {
        return num
    }
    return Fibonnaci(num-1) + Fibonnaci(num-2)
}

func worker(id int, cjobs <-chan int, cresults chan<- int) {
    for job := range cjobs {
        fmt.Printf("worker %d start fibonacci for %d\n", id, job)
        fib := Fibonnaci(job)
        fmt.Printf("Worker with id %d, job %d, and fibonacci %d\n", id, job, fib)
        cresults <- fib
    }
}

func main() {
    taks := []int {2,3,4,5,7,10,12,40}
    nWorker := 3
    cjobs := make(chan int, len(taks))
    cresults := make(chan int, len(taks))

    for i:=0; i<nWorker; i++ {
        go worker(i, cjobs, cresults) 
    }

    for _, value := range taks {
        cjobs <- value
    }
    close(cjobs)

    for r:=0; r<len(taks); r++ {
        <-cresults
    }
}
