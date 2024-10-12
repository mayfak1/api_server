package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", handleHelloworld)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/new-endpoint", handleNewEndpoint)
	fmt.Println("added print statement")
	addr := "localhost:8000"
	log.Printf("Lostening on %s ...", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
func handleHelloworld(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	write(writer, "Hello from func")

}
func handleHealth(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	write(writer, "hello from health")
}
func handleNewEndpoint(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	write(writer, "hello from new-endpoint")
}
func write(writer http.ResponseWriter, msg string) {
	response := []byte(msg)
	fmt.Println(response)

	_, err := writer.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}
