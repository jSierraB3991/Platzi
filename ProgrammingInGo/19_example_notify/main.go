package main

import (
   notify "github.com/mqu/go-notify" 
    "os"
    "fmt"
    "time"
)


const (
        DELAY = 3000;
    )


func main() {
    notify.Init("Hello world")
    hello := notify.NotificationNew("Hello World!",
            "This is an example notification.",
                    "")
    if hello == nil {
        fmt.Fprintf(os.Stderr, "Unable to create a new notification\n")
        return
    }

    notify.NotificationSetTimeout(hello, DELAY)

    if e := notify.NotificationShow(hello); e != nil {
        fmt.Fprintf(os.Stderr, "%s\n", e.Message())
        return
    }

    time.Sleep(DELAY * 1000000)
    notify.NotificationClose(hello)

    notify.UnInit()

    
}
