package server

import (
	"net/http"

	"github.com/sullyh7/shahid/internal/store"
	"github.com/sullyh7/shahid/view/home"
)

func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch videos from the store
	videos := []store.Video{
		{
			ID:          1,
			Title:       "Building a Go Web App",
			Description: "Learn how to build a web app in Go using Templ and Tailwind.",
			Thumbnail:   "https://placehold.co/600x400/EEE/31343C",
			URL:         "/watch/1",
			Duration:    732,
			UploadedAt:  "2025-06-01T14:00:00Z",
			ChannelID:   101,
			ChannelName: "Go Dev Hub",
		},
		{
			ID:          2,
			Title:       "Deploying to a VPS",
			Description: "A practical guide to deploying full stack apps to a VPS.",
			Thumbnail:   "https://placehold.co/600x400/EEE/31343C",
			URL:         "/watch/2",
			Duration:    1250,
			UploadedAt:  "2025-05-15T10:30:00Z",
			ChannelID:   102,
			ChannelName: "DevOps Digest",
		},
		{
			ID:          3,
			Title:       "TailwindCSS in 15 Minutes",
			Description: "Master the basics of TailwindCSS quickly and efficiently.",
			Thumbnail:   "https://placehold.co/600x400/EEE/31343C",
			URL:         "/watch/3",
			Duration:    905,
			UploadedAt:  "2025-04-10T08:00:00Z",
			ChannelID:   103,
			ChannelName: "Frontend Brief",
		},
		{
			ID:          4,
			Title:       "Understanding HTTP in Go",
			Description: "Dig into net/http, request handling, and middleware.",
			Thumbnail:   "https://placehold.co/600x400/EEE/31343C",
			URL:         "/watch/4",
			Duration:    1180,
			UploadedAt:  "2025-03-21T16:45:00Z",
			ChannelID:   101,
			ChannelName: "Go Dev Hub",
		},
		{
			ID:          5,
			Title:       "Intro to Templ",
			Description: "A beginner's walkthrough of building interfaces with Templ.",
			Thumbnail:   "https://placehold.co/600x400/EEE/31343C",
			URL:         "/watch/5",
			Duration:    670,
			UploadedAt:  "2025-06-10T12:00:00Z",
			ChannelID:   104,
			ChannelName: "UI Toolkit",
		},
	}

	if err := home.BrowsePage(videos).Render(r.Context(), w); err != nil {
		s.InternalServerError(w, r, err)
		return
	}

}
