package main

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"github.com/teten-nugraha/go-react/models"
	"net/http"
	"strconv"
	"time"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("Invalid id parameter"))
	}

	app.logger.Println("id is ", id)

	movie := models.Movie{
		ID:          id,
		Title:       "Some movie",
		Description: "Desc",
		Year:        2021,
		ReleaseDate: time.Date(2021,01,01,01,0,0,0, time.Local),
		Runtime:     0,
		Rating:      5,
		MPAARating:  "PG-13",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {

}