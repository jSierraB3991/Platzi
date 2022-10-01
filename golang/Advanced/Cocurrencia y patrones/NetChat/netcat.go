package main

import(
    "flag"
    "fmt"
    "io"
    "log"
    "net"
    "os"
)

func CopyContent(dst io.Writer, src io.Reader) {
    _, err := io.Copy(dst, src)
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    port := flag.Int("p", 3090, "Port for scann")
    host := flag.String("h", "localhost", "host for scann")
    flag.Parse()

    conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
    if err != nil {
        log.Fatal(err)
    }

    done := make(chan struct{})

    go func() {
        io.Copy(os.Stdout, conn)
        done <- struct{}{}
    }()
    CopyContent(conn, os.Stdin)

    conn.Close()
    <- done
}
