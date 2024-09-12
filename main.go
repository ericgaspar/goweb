package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(10000000)
	fmt.Fprintf(w, "Random Number: %d", randomNumber)
}
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
