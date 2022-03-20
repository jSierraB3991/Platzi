package main

import (
    "fmt"
    "net/http"
    "time"
)

func review_server(server string, channel chan string) {
    uri := "http://" + server 
    _, err := http.Get(uri)
    if err != nil {
        channel <- uri + " no available :("
    }else {
        channel <- uri + " is available :)"
    }
}

func main() {
    start := time.Now()
    channel := make(chan string)
    servers := []string { "platzi.com", "google.com", "facebook.com", "instagram.com" }
    
    for _, value := range servers {
        go review_server(value, channel)
    }
    for i := 0; i < len(servers); i++ {
        fmt.Println(<-channel)
    }
    tiemp_pass := time.Since(start)
    fmt.Printf("Pass tiemp %s\n", tiemp_pass)
}
