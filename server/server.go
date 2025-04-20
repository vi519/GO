package main

import (
	"fmt"
	"net/http"
)

func myMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, You are god damm right")

}

func main() {
	http.HandleFunc("/SayMyName=Vineet", myMessage)
	fmt.Println("Server is running bro 8181")
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		println("Server is down")
	}
}
