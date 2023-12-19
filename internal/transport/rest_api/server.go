package rest_api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"

	"github.com/hathawway/back-gloss/internal/config"
)

type Server struct {
	HttpServer *http.Server

	version string
}

func NewServer(cfg *config.Config) *Server {
	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	s := &Server{
		HttpServer: &http.Server{
			Addr:    "0.0.0.0:" + strconv.Itoa(cfg.GetInt(config.ServerRestAPIPort)),
			Handler: c.Handler(r),
		},

		version: cfg.GetString(config.AppInfoVersion),
	}

	r.HandleFunc("/api/posts", s.Posts)
	r.HandleFunc("/api/search", s.Search)
	r.HandleFunc("/api/graph", s.Graph)
	return s
}

func (s *Server) Start(ctx context.Context) error {
	go func() {
		err := s.HttpServer.ListenAndServe()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			logrus.Fatal(err)
		}
	}()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.HttpServer.Shutdown(ctx)
}

func (s *Server) formResponse(r interface{}) ([]byte, error) {
	return json.Marshal(r)
}
