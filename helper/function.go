package helper

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// GetMovies
func GetMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		movies := Movies()
		dataMovies, err := json.Marshal(movies)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataMovies)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

var movs = []Movie{
	{1, "Spider-Man", 2002},
	{2, "John Wick", 2014},
	{3, "Avengers", 2018},
	{4, "Logan", 2017},
}

func Movies() []Movie {
	return movs
}

// PostMovie
func PostMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Mov Movie
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Mov); err != nil {
				log.Fatal(err)
			}
		} else {
			// parse dari form
			getID := r.PostFormValue("id")
			id, _ := strconv.Atoi(getID)
			title := r.PostFormValue("title")
			getYear := r.PostFormValue("year")
			year, _ := strconv.Atoi(getYear)
			Mov = Movie{
				ID:    id,
				Title: title,
				Year:  year,
			}
		}
		movs = append(movs, Mov)
		dataMovie, _ := json.Marshal(Mov) // to byte
		w.Write(dataMovie)
		return
	}

	http.Error(w, "NOT FOUND", http.StatusNotFound)
	return
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usernaame, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Mandatory authentication required"))
			return
		}
		if usernaame == "admin" && password == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		w.Write([]byte("Invalid username/password"))
		return
	})
}
