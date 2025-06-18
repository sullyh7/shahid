package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Video struct {
	ID            string
	ChannelID     string
	Title         string
	ThumbnailPath string
	VideoPath     string
	Status        string
	ErrorMessage  *string
	CreatedAt     time.Time
	ProcessedAt   *time.Time
}

var ErrVideoAlreadyExists = errors.New("video already exists")

type VideoStore struct {
	db *sql.DB
}

// make a new video entry (unprocessed)
func (s *VideoStore) Create(ctx context.Context, video *Video) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancel()

	query := `
		INSERT INTO videos (id, channel_id, title, thumbnail_path, video_path, status, error_message)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := s.db.ExecContext(ctx, query,
		video.ID,
		video.ChannelID,
		video.Title,
		video.ThumbnailPath,
		video.VideoPath,
		video.Status,
		video.ErrorMessage,
	)

	if isUniqueViolation(err) {
		return ErrVideoAlreadyExists
	}

	return err
}
