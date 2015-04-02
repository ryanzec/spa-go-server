package main

import (
  "github.com/codegangsta/negroni"
  "net/http"
  "fmt"
  "flag"
  "strconv"
)

var (
  port int
)

func main() {
  startServer()
}

func startServer() {
  flag.IntVar(&port, "port", 3000, "The port to run the server on")

  flag.Parse();

  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page!")
  })

  n := negroni.Classic()
  n.UseHandler(mux)
  n.Run(":" + strconv.Itoa(port))
}