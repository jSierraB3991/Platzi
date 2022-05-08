package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "net"
)

type Client chan<- string

var (
    incommingClient = make(chan Client)
    leavingClient   = make(chan Client)
    messages        = make(chan string)
)

var (
    host = flag.String("h", "localhost", "host for server")
    port = flag.Int("p", 3090, "port for server")
)

func HandleConecction(conn net.Conn){
    defer conn.Close()
    messageClient := make(chan string)
    go MessageWrite(conn, messageClient)

    clientName  := conn.RemoteAddr().String()
    messageClient<-fmt.Sprintf("Welcome to server %s", clientName)
    messages <- fmt.Sprintf("New client here is, Client: %s", clientName)

    incommingClient <- messageClient

    inputMessage := bufio.NewScanner(conn)
    for inputMessage.Scan() {
        messages <- fmt.Sprintf("%s: %s", clientName, inputMessage.Text())
    }
    leavingClient <- messageClient
    messages<- fmt.Sprintf("Client %s say Good bye!", clientName)
}

func MessageWrite(conn net.Conn, messageClient <-chan string){
    for message := range messageClient {
        fmt.Fprintln(conn, message)
    }
}

func Broadcast() {
    clients := make(map[Client]bool)
    for {
        select {
            case message := <-messages:
                for client := range clients {
                    client <- message
                }
            case newClient := <-incommingClient:
                clients[newClient] = true
            case removeClient := <-leavingClient:
                delete(clients, removeClient)
                close(removeClient)
        }
    }
}

func main() {
    flag.Parse()

    listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
    if err != nil {
        log.Fatal(err)
    }
    go Broadcast()
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Print(err)
            continue
        }
        go HandleConecction(conn)
    }
}
