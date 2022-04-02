package main

func main() {
    server := NewServer(":5000")
    server.Handle("GET", "/", HandlerRoot)
    server.Handle("POST", "/api", server.AddMiddleware(HandlerHome, CheckAuth(), Logging()))
    server.Handle("POST", "/create", PostRequest)
    server.Listen()
}
