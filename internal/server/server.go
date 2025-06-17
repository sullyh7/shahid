package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sullyh7/shahid/assets"
	"github.com/sullyh7/shahid/internal/store"
	"github.com/sullyh7/shahid/internal/worker"

	"go.uber.org/zap"
)

type Server struct {
	Config Config
	Store  store.Storage
	Logger *zap.SugaredLogger
	Queue  *worker.WorkerPool
}

type Config struct {
	Addr    string
	Db      DBConfig
	Env     string
	Version string
}

type DBConfig struct {
	Addr         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

func (s *Server) Mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	var isDevelopment = s.Config.Env != "production"

	assetHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isDevelopment {
			w.Header().Set("Cache-Control", "no-store")
		}

		var fs http.Handler
		if isDevelopment {
			fs = http.FileServer(http.Dir("./assets"))
		} else {
			fs = http.FileServer(http.FS(assets.Assets))
		}

		fs.ServeHTTP(w, r)
	})

	r.Handle("/assets/*", http.StripPrefix("/assets/", assetHandler))

	r.Post("/callback", s.WebhookCallback)
	r.Get("/", s.HomeHandler)

	return r
}

func (s *Server) Run(mux http.Handler) error {
	srv := http.Server{
		Addr:         s.Config.Addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	s.Logger.Infow("server has started", "addr", s.Config.Addr, "env", s.Config.Env)
	return srv.ListenAndServe()
}
