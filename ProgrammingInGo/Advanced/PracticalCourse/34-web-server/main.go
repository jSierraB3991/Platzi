package main

func main() {
    server := NewServer(":5000")
    server.Handle("/", HandlerRoot)
    server.Handle("/api", HandlerHome)
    server.Listen()
}
