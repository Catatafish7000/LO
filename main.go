package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cache := NewCache()
	db, err := NewPsqlDB(ConfigDB{
		Host:     "localhost",
		Port:     5432,
		Username: "bob",
		DBName:   "test-db",
		Password: "admin",
	})

	if err != nil {
		log.Fatalf("error creating database: %s", err.Error())
	}

	ex := Exchange{
		C:  cache,
		DB: db,
	}

	err = ex.GetCache()
	if err != nil {
		log.Fatalf("error initializing cache: %s", err.Error())
	}

	natsCfg := MsgProcessor{
		StanClusterID: "test-cluster",
		ClientID:      "client-1000",
		SubjName:      "foo",
		QGroup:        "group",
		DurableName:   "my-durable",
		cache:         &ex,
		SubsAmount:    3,
	}

	srv := New(":8080", cache)
	go func() {
		if err := srv.Serve(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error running http server: %v", err)
		}
	}()
	err = natsCfg.Connect(natsCfg.MessageHandler())
	if err != nil {
		log.Fatalf("error running nats-streaming server: %v", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err = srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}
	if err = db.Close(); err != nil {
		log.Printf("error occured on cache connection close: %s", err.Error())
	}
}
