package main

import (
    "bufio"
    "log"
    "fmt"
    "os"
)

var path = "/home/juan-sierra/Documents/file.txt"

func haveError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }
    return (err != nil)
}

func writeFile() {
    var file, err = os.OpenFile(path, os.O_RDWR, 0644)
    if haveError(err) {
        return
    }
    defer file.Close()
    _, err = file.WriteString("Hello \n")
    if haveError(err) {
        return
    }
    _, err = file.WriteString("World \n")
    if haveError(err) {
        return
    }
    err = file.Sync()
    if haveError(err) {
        return
    }
    fmt.Println("File update success.")
}

func createFile() {
    var _, err = os.Stat(path)
    if os.IsNotExist(err) {
        var file, err = os.Create(path)
        if haveError(err) {
            return
        }
        defer file.Close()
    }
    fmt.Println("File Created Successfully", path)
}

func readFile() {
    file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}

func main() {
    createFile()
    writeFile()
    readFile()
}
