package main

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

type User struct {
    Name string
    Email string
    Phone string
}

type Metadata interface {}
