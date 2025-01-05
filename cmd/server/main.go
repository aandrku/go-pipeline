package main

import (
	"fmt"
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

	http.ListenAndServe(":8888", mux)
}
