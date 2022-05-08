package main

import (
    "time"
    "fmt"
    "net/http"
    "strconv"
    "log"
)

type Job struct {
    Name string
    delay time.Duration
    Number int
}

type Worker struct {
    Id int
    JobQueue chan Job
    WorkerPool chan chan Job
    QuitChan chan bool
}

type Dispatcher struct {
    WorkerPool chan chan Job
    MaxWorkers int
    JobQueue chan Job
}

func NewWorker(id int, workerPool chan chan Job) *Worker {
    return &Worker {
        Id: id,
        WorkerPool: workerPool,
        JobQueue: make(chan Job),
        QuitChan: make(chan bool),
    }
}

func (w Worker) Start(){
    go func(){
        for {
            w.WorkerPool <- w.JobQueue
            select {
            case job := <- w.JobQueue:
                fmt.Printf("Worker with id %d Started\n", w.Id)
                fib := Fibonnaci(job.Number)
                time.Sleep(job.delay)
                fmt.Printf("Worker with id %d finally with fibonnaci number %d \n", w.Id, fib)

            case <- w.QuitChan:
                fmt.Printf("Worker with id %d is finnalize\n", w.Id)
            }
        }
    }()
}

func (w *Worker) Stop() {
    go func () {
        w.QuitChan <- true
    }()
}

func Fibonnaci(num int) int {
    if num <= 1 {
        return num 
    }
    return Fibonnaci(num-1)+Fibonnaci(num-2)
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) Dispatcher {
    return Dispatcher {
        WorkerPool: make(chan chan Job, maxWorkers),
        MaxWorkers: maxWorkers,
        JobQueue: jobQueue,
    }
}

func (d *Dispatcher) Dispatch() {
    for {
        select {
        case job := <-d.JobQueue:
            go func(){
                workerJobQueue := <- d.WorkerPool
                workerJobQueue <- job
            }()
        }
    }
}

func (d *Dispatcher) Run() {
    for i:=0; i<d.MaxWorkers; i++ {
        worker := NewWorker(i, d.WorkerPool)
        worker.Start()
    }

    go d.Dispatch()
}

func RequestHandler(w http.ResponseWriter,request *http.Request, jobQueue chan Job) {
    if request.Method != "POST" {
        w.Header().Set("Allow", "POST")
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    delay, err := time.ParseDuration(request.FormValue("delay"))
    if err != nil {
        http.Error(w, "Invalid delay", http.StatusBadRequest)
        return
    }

    value, err := strconv.Atoi(request.FormValue("value"))
    if err != nil {
        http.Error(w, "Invalid value", http.StatusBadRequest)
        return
    }

    name := request.FormValue("name")
    if name == "" {
        http.Error(w, "Invalid name", http.StatusBadRequest)
        return
    }
    job := Job{Name: name, delay: delay, Number: value,}
    jobQueue <- job
    w.WriteHeader(http.StatusCreated)
}


func main(){
    const (
        maxWorkers = 4
        maxJobs = 20
        port = ":8081"
    )

    jobQueue := make(chan Job, maxJobs)
    dispatcher := NewDispatcher(jobQueue, maxWorkers)
    dispatcher.Run()

    // http://localhost:8081/fib
    http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
        RequestHandler(w, r, jobQueue)
    })
    log.Fatal(http.ListenAndServe(port, nil))
}
