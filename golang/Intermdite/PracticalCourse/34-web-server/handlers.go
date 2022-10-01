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
    var user Metadata
    err := decoder.Decode(&user)
    if err != nil {
        fmt.Fprintf(w, "error: %v\n", err)
        return
    }
    fmt.Fprintf(w, "Payload %v\n", user)
}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var user User
    err := decoder.Decode(&user)
    if err != nil {
        fmt.Fprintf(w, "error: %v\n", err)
        return
    }
    fmt.Println(user.Name)
    response, err := user.ToJson()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    //w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}
