package helper

import (
	"encoding/json"
	"net/http"
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

func Movies() []Movie {
	movs := []Movie{
		{1, "Spider-Man", 2002},
		{2, "John Wick", 2014},
		{3, "Avengers", 2018},
		{4, "Logan", 2017},
	}
	return movs
}
