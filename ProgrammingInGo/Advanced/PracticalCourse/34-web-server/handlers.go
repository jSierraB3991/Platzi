package main

import (
    "net/http"
    "fmt"
)

func handlerRoot(w http.ResposeWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world")
}
