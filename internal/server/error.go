package server

import (
	"encoding/json"
	"log"
	"net/http"

	errorview "github.com/sullyh7/shahid/view/error"
)

func (s *Server) InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s err: %s", r.Method, r.URL.Path, err)
	s.Logger.Errorw("internal server error",
		"path", r.URL.Path,
		"method", r.Method,
		"error", err.Error(),
	)
	writeHTMLError(w, r, http.StatusInternalServerError, "There was a problem.")
}

func (s *Server) NotFound(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found: %s path: %s err: %s", r.Method, r.URL.Path, err)
	s.Logger.Warnw("not found",
		"path", r.URL.Path,
		"method", r.Method,
		"error", err.Error(),
	)
	writeHTMLError(w, r, http.StatusNotFound, "The requested resource could not be found.")
}

func (s *Server) BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request: %s path: %s err: %s", r.Method, r.URL.Path, err)
	s.Logger.Warnw("bad request",
		"path", r.URL.Path,
		"method", r.Method,
		"error", err.Error(),
	)
	writeHTMLError(w, r, http.StatusBadRequest, "Bad request: "+err.Error())
}

func (s *Server) Unauthorized(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("unauthorized: %s path: %s", r.Method, r.URL.Path)
	s.Logger.Warnw("unauthorized",
		"path", r.URL.Path,
		"method", r.Method,
		"err", err.Error(),
	)
	writeHTMLError(w, r, http.StatusUnauthorized, "You are not authorized to access this resource.")
}

func (s *Server) Forbidden(w http.ResponseWriter, r *http.Request) {
	log.Printf("forbidden: %s path: %s", r.Method, r.URL.Path)
	s.Logger.Warnw("forbidden",
		"path", r.URL.Path,
		"method", r.Method,
	)
	writeHTMLError(w, r, http.StatusForbidden, "You do not have permission to access this resource.")
}

func (s *Server) Conflict(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("conflict: %s path: %s err: %s", r.Method, r.URL.Path, err)
	s.Logger.Warnw("conflict",
		"path", r.URL.Path,
		"method", r.Method,
		"error", err.Error(),
	)
	writeHTMLError(w, r, http.StatusConflict, "Conflict: "+err.Error())
}

func writeHTMLError(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	err := errorview.ErrorPage(message).Render(r.Context(), w)
	if err != nil {
		log.Printf("templ render failed: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// may need at some point
func readJson(w http.ResponseWriter, r *http.Request, data any) error {
	maxByes := 1_048_578
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxByes))
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(data)
}
