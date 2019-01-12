package main

import (
  "net/http"
  "fmt"
  "log"
)

func main() {
  http.HandleFunc("/", Handler)
  log.Fatal(http.ListenAndServe(":5000", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "hello")
}
