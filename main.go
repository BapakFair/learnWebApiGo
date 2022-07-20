package main

import (
	"fmt"
	. "learn-web-api/helper"
	"log"
	"net/http"
)

func main() {
	http.Handle("/movies", Auth(http.HandlerFunc(GetMovies)))
	http.Handle("/post_movie", Auth(http.HandlerFunc(PostMovie)))

	fmt.Println("server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
