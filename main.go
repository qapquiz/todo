package main

import (
  "net/http"
  "fmt"
  "log"
  "os"
)

func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}

func main() {
  http.HandleFunc("/", Handler)

  address, err := determineListenAddress()
  if err != nil{ 
    log.Fatal("$PORT not set")
  }

  log.Fatal(http.ListenAndServe(address, nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "hello")
}
