package server

import (
	"encoding/json"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"reflect"
)

type server struct {
	logger   *zap.Logger
	router   gin.IRouter
	esClient *esdb.Client
}

func New(l *zap.Logger, r gin.IRouter, c *esdb.Client) *server {
	return &server{l, r, c}
}

func (s *server) decodeRequestFails(ctx *gin.Context, req interface{}) bool {
	err := ctx.BindJSON(req)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Status(http.StatusBadRequest)
	}
	return err != nil
}

func (s *server) encodeEvent(obj interface{}, eventType string) esdb.EventData {
	data, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	return esdb.EventData{
		EventType:   eventType,
		ContentType: esdb.JsonContentType,
		Data:        data,
	}
}

func (s *server) appendToStreamFails(ctx *gin.Context, streamId string, expected esdb.ExpectedRevision, evt interface{}) bool {
	et := reflect.TypeOf(evt).Elem().Name()
	ed := s.encodeEvent(evt, et)
	opts := esdb.AppendToStreamOptions{ExpectedRevision: expected}
	_, err := s.esClient.AppendToStream(ctx, streamId, opts, ed)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	}
	return err != nil
}
