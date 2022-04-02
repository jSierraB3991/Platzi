package main

import (
    "fmt"
    "log"
    "time"
    "net/http"
)

func CheckAuth() Middleware {
    return func(f http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            flag := true
            fmt.Println("Checking Authentication")
            if flag {
                f(w, r)
            }else {
                w.WriteHeader(http.StatusUnauthorized)
            }
        }
    }
}

func Logging() Middleware {
    return func(f http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            defer func() {
                log.Println(r.URL.Path, time.Since(start))
            }()
            f(w, r)
        }
    }
}
