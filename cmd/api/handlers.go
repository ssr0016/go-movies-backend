package main

import (
	"errors"
	"log"
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
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, movies)
}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	// read json payload
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate user against database
	user, err := app.DB.GetuserByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// check password

	// create a jwt user
	u := jwtUser{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
	}

	// generate tokens
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	log.Println(tokens.Token)
	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	w.Write([]byte(tokens.Token))
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
