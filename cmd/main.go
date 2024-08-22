package main

import (
    "github.com/Damianen/appie-app/handler"
    "fmt"
    "log"
    "net/http"
)

func main() {
    loginHandler := handler.LoginHandler{}

    fileServer := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fileServer))

    http.HandleFunc("/", loginHandler.HandleLoginShow)

    fmt.Println("Server is listening on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
