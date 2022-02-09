package server

import (
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	events "github.com/tyronegroves/cqrs-sample-events"
	"net/http"
)

func (s *server) handleReleaseMovie() gin.HandlerFunc {
	type request struct {
		Title     string   `json:"title"`
		PosterUrl string   `json:"poster_url"`
		Genres    []string `json:"genres"`
		Runtime   int      `json:"runtime"`
		Rating    string   `json:"rating"`
	}

	type response struct {
		MovieId uuid.UUID `json:"movie_id"`
	}

	return func(ctx *gin.Context) {
		req := &request{}
		if s.decodeRequestFails(ctx, req) {
			return
		}

		movieId := uuid.Must(uuid.NewV4())
		evt := &events.MovieReleased{
			MovieId:   movieId,
			Title:     req.Title,
			PosterUrl: req.PosterUrl,
			Genres:    req.Genres,
			Runtime:   req.Runtime,
			Rating:    req.Rating,
		}

		if s.appendToStreamFails(ctx, "movie-"+movieId.String(), esdb.NoStream{}, evt) {
			return
		}

		ctx.JSON(http.StatusCreated, &response{movieId})
	}
}
