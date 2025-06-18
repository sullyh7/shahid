package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sullyh7/shahid/internal/db"
	"github.com/sullyh7/shahid/internal/env"
	"github.com/sullyh7/shahid/internal/server"
	"github.com/sullyh7/shahid/internal/store"
	videoservice "github.com/sullyh7/shahid/internal/video_service"
	"github.com/sullyh7/shahid/internal/worker"
	"go.uber.org/zap"
)

func main() {
	cfg := server.Config{
		Addr: env.GetString("ADDR", ":8080"),
		Db: server.DBConfig{
			Addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/shahid?sslmode=disable"),
			MaxOpenConns: env.GetInt("MAX_OPEN_CONNS", 30),
			MaxIdleConns: env.GetInt("MAX_OPEN_CONNS", 30),
			MaxIdleTime:  env.GetString("MAX_IDLE_TIME", "15m"),
		},
		Env:     env.GetString("ENV", "development"),
		Version: "1",
	}

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	db, err := db.New(cfg.Db.Addr, cfg.Db.MaxOpenConns, cfg.Db.MaxIdleConns, cfg.Db.MaxIdleTime)
	if err != nil {
		logger.Fatalw("error connecting to db", "err", err)
	}
	defer db.Close()
	store := store.NewStorage(db)

	vs := &videoservice.VideoService{
		Store:  store,
		FS:     &videoservice.LocalFileService{},
		TS:     &videoservice.StandardTranscriptService{},
		Logger: logger,
	}

	wp := &worker.WorkerPool{
		JobChan: make(chan worker.Job, 5),
		Size:    10,
		Process: vs.Process,
		Logger:  logger,
	}

	server := &server.Server{
		Config: cfg,
		Store:  store,
		Logger: logger,
		Queue:  wp,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go wp.Run(ctx)

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		logger.Infow("shutdown signal received")
		cancel()
	}()

	mux := server.Mount()
	if err = server.Run(mux); err != nil {
		logger.Fatalw(err.Error())
	}
}
