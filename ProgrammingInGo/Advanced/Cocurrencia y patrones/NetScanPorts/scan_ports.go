package main

import (
    "flag"
    "fmt"
    "net"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    site := flag.String("site", "scanme.nmap.org", "Url for scann ports")
    
    flag.Parse()
    fmt.Printf("Scann port for site %s\n", *site)
    for i:=0; i < 65535; i++ {
        wg.Add(1)
        go func(port int) {
            defer wg.Done()

            conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
            if err != nil {
                return
            }
            conn.Close()
            fmt.Printf("The port %d is Open\n", port)
        }(i)
    }
    wg.Wait()
}
