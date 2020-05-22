package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":6060", nil))
}

func main() {
    s := getSomeStruct()
    fmt.Println(s.name)
    handleRequests()
}

type SomeStruct struct {
	name string
}


func getSomeStruct() *SomeStruct {
	return nil
}