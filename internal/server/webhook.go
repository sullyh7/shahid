package server

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/sullyh7/shahid/internal/store"
	"github.com/sullyh7/shahid/internal/worker"
)

type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Entry   Entry    `xml:"entry"`
}

type Entry struct {
	VideoID   string `xml:"http://www.youtube.com/xml/schemas/2015 videoId"`
	ChannelID string `xml:"http://www.youtube.com/xml/schemas/2015 channelId"`
	Title     string `xml:"title"`
}

func (s *Server) WebhookCallback(w http.ResponseWriter, r *http.Request) {
	var feed Feed
	if err := xml.NewDecoder(r.Body).Decode(&feed); err != nil {
		s.BadRequest(w, r, err)
		return
	}
	defer r.Body.Close()

	if feed.Entry.VideoID == "" {
		s.BadRequest(w, r, fmt.Errorf("video ID not found in feed"))
		return
	}
	if feed.Entry.ChannelID == "" {
		s.BadRequest(w, r, fmt.Errorf("channel ID not found in feed"))
		return
	}

	if feed.Entry.Title == "" {
		s.BadRequest(w, r, fmt.Errorf("no title to video"))
		return
	}
	s.Logger.Infof("Received webhook for video ID: %s, Channel ID: %s", feed.Entry.VideoID, feed.Entry.ChannelID)

	video := &store.Video{
		ID:     feed.Entry.VideoID,
		Status: "processing",
	}

	if err := s.Store.Videos.Create(r.Context(), video); err != nil {
		s.InternalServerError(w, r, err)
		return
	}

	if err := s.Queue.Add(worker.Job{
		VideoID: feed.Entry.VideoID,
	}); err != nil {
		s.InternalServerError(w, r, err)
	}
	w.WriteHeader(http.StatusNoContent)
}
