package main

 

import (

  "fmt"

  "log"

  "net/http"

)

func index(w http.ResponseWriter, r *http.Request) {

  fmt.Fprintf(w, "Halo ini docker golang testsss")

}

 

func about(w http.ResponseWriter, r *http.Request) {

  fmt.Fprintf(w, "Halo ini test")

}

func main() {

  http.HandleFunc("/", index)

  http.HandleFunc("/about", about)

 

  log.Println("application started at port :8990")

  log.Fatal(http.ListenAndServe(":8990", nil))

}
