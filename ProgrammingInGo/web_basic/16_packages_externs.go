package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":1323"))
}


// installs og go get
// 1) go mod init {repo:github.com}/{user_name}/{name_proyect}
// 2) go get {package} // example: go get github.com/labstack/echo/v4
//  add get {-v} package, for information to download
//  add get {-u} package, for download package, another
// 3) go run server.go
