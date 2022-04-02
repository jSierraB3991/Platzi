package main

import (
    "net/http"
    "fmt"
)

func HandlerRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world")
}

func HandlerHome(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is th API Enpoint")
}
