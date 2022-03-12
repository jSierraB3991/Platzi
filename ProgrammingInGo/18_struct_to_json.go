package main

import (
    "encoding/json"
    "fmt"
)

type UserS struct {
    Name   string
    UserP  string
    Password string
    Email    string
}

func main() {
    user := &UserS{Name: "Eliot", UserP: "Musk", Password: "123", Email: "xxx@prueba.com"}
    userJSON, err := json.Marshal(user)
    if err != nil {
        fmt.Printf("Error: %s", err)
        return
    }
    fmt.Println(string(userJSON))
}
