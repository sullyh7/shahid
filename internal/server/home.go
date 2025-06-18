package server

import (
	"net/http"

	"github.com/sullyh7/shahid/view/home"
)

func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch videos from the store

	if err := home.BrowsePage(nil).Render(r.Context(), w); err != nil {
		s.InternalServerError(w, r, err)
		return
	}

}
