package main

import (
	"fmt"
	. "learn-web-api/helper"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/movies", GetMovies)
	http.HandleFunc("/post_movie", PostMovie)

	fmt.Println("server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
