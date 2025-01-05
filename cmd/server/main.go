package main

import (
	"fmt"
	"log"
	"net/http"
)

type SomeStruct struct {
	Name   string
	Job    string
	FooBar string
}

func main() {
	fmt.Println("Hello, world")

	mux := http.NewServeMux()

	if err := http.ListenAndServe(":8888", mux); err != nil {
		log.Fatal("Failed to start a server")
	}
}
