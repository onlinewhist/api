package api

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/api/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}
