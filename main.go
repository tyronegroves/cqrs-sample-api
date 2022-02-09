package main

import (
	"github.com/EventStore/EventStore-Client-Go/esdb"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/tyronegroves/cqrs-sample-api/config"
	"github.com/tyronegroves/cqrs-sample-api/server"
	"go.uber.org/zap"
	"log"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg, err := config.LoadConfiguration()
	if err != nil {
		return err
	}

	l, err := zap.NewDevelopment()
	if err != nil {
		return err
	}

	settings, err := esdb.ParseConnectionString(cfg.EventStoreDb.Url)
	if err != nil {
		return err
	}

	db, err := esdb.NewClient(settings)
	if err != nil {
		return err
	}
	defer db.Close()

	r := gin.New()
	r.Use(ginzap.Ginzap(l, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(l, true))

	if err := r.SetTrustedProxies(cfg.TrustedProxies); err != nil {
		return err
	}

	s := server.New(l, r, db)
	s.LoadRoutes()

	if err := r.Run(cfg.LocalAddresses...); err != nil {
		return err
	}
	return nil
}
