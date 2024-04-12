package main

import (
	"fmt"
	"net/http"
	"templates_study/renderfunc"
)

func main() {
	http.HandleFunc("/", renderfunc.SayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP server failed: %v\n", err)
		return
	}
}
