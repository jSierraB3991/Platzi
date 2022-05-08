package main

import (
    "fmt"
    "sync"
    "time"
)

type Database struct {}

var db *Database
var mutex sync.Mutex

func (Database) CreateSingleConnection() {
    fmt.Println("Create singleton for database")

}

func GetDataBaseConnection() *Database {
    mutex.Lock()
    if db == nil {
        fmt.Println("Create Database Connecton")
        db = &Database{}
        time.Sleep(2 * time.Second)
        db.CreateSingleConnection()
    } else {
        fmt.Println("DB realy created")
    }
    mutex.Unlock()
    return db
}

func main() {
    var wg sync.WaitGroup
    
    for i:=0; i<10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            GetDataBaseConnection()
        }()
    }
    wg.Wait()
}
