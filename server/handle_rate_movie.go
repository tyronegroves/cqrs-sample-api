package server

import (
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	events "github.com/tyronegroves/cqrs-sample-events"
	"net/http"
)

func (s *server) handleRateMovie() gin.HandlerFunc {
	type request struct {
		MovieId  string `json:"movieId"`
		Username string `json:"username"`
		Rating   int    `json:"rating"`
	}

	type response struct {
		RatingId uuid.UUID `json:"ratingId"`
	}

	return func(ctx *gin.Context) {
		req := &request{}
		if s.decodeRequestFails(ctx, req) {
			return
		}

		ratingId := uuid.Must(uuid.NewV4())
		evt := &events.MovieRated{
			RatingId: ratingId,
			MovieId:  req.MovieId,
			Username: s.username(ctx),
			Rating:   req.Rating,
		}

		if s.appendToStreamFails(ctx, "rating-"+ratingId.String(), esdb.NoStream{}, evt) {
			return
		}

		ctx.JSON(http.StatusOK, &response{ratingId})
	}
}
