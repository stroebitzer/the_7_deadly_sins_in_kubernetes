package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var memLeak [][]byte

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("hello called")
	fmt.Fprintf(w, "Hello!")
}

func memLeakHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("memLeak called")
	fmt.Fprintf(w, "Memory Leak")
	createMemLeak()
}

func createMemLeak() {
	for i := 0; i < 10; i++ {
		// ~ 30 MB
		memLeak = append(memLeak, make([]byte, 1024*1024*1024))
		time.Sleep(1 * time.Second)
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/memleak", memLeakHandler)

	isMemLeak := os.Getenv("MEM_LEAK")

	if isMemLeak == "true" {
		log.Println("In memory leak mode")
		waitSeconds := rand.Intn(60)
		time.Sleep(time.Duration(waitSeconds) * time.Second)
		log.Printf("Waiting %d seconds", waitSeconds)
		createMemLeak()
	}

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
