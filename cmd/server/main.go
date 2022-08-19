package main

import (
	"net/http"
)


func main() {
	err := http.ListenAndServe(":5050", http.FileServer(http.Dir("../../assets")))
	if err != nil {
		panic(err)
	}
}
