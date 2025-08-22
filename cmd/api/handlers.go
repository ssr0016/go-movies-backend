package main

import (
	"fmt"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Movies up  and running",
		Version: "1.0.0",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()
	if err != nil {
		fmt.Println(err)
	}

	_ = app.writeJSON(w, http.StatusOK, movies)
}

// highlander := models.Movie{
// 	ID:          1,
// 	Title:       "Highlander",
// 	ReleaseDate: rd,
// 	MPAARating:  "R",
// 	RunTime:     116,
// 	Description: "In the year 1536, a Scottish Highlander named Connor MacLeod is mortally wounded in battle, but mysteriously recovers.",
// 	CreatedAt:   time.Now(),
// 	UpdatedAt:   time.Now(),
// }

// rotla := models.Movie{
// 	ID:          2,
// 	Title:       "Raiders of the Lost Ark",
// 	ReleaseDate: rds,
// 	MPAARating:  "PG-13",
// 	RunTime:     115,
// 	Description: "Another classic adventure film featuring Indiana Jones.",
// 	CreatedAt:   time.Now(),
// 	UpdatedAt:   time.Now(),
// }
