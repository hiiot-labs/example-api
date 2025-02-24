package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello world"))
	})
	http.HandleFunc("/test/v2", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello v2"))
	})

	fmt.Println("Listening on port 8899")
	err := http.ListenAndServe(":8899", nil)
	if err != nil {
		panic(err)
	}
}
