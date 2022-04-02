package main

import (
    "net/http"
    "fmt"
    "encoding/json"
)

func HandlerRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world")
}

func HandlerHome(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is th API Enpoint")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var metadata Metadata
    err := decoder.Decode(&metadata)
    if err != nil {
        fmt.Fprintf(w, "error: %v\n", err)
        return
    }
    fmt.Fprintf(w, "Payload %v\n", metadata)
}
