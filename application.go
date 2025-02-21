package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, Lucile!")
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "5000"
    }

    fmt.Printf("Listening on port %s...\n", port)
    http.ListenAndServe(":"+port, nil)
}
