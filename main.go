package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "im iron man!")
    })

    fmt.Println("Server running on port 8082")
    http.ListenAndServe(":8082", nil)
}