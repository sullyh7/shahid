package videoservice

import (
	"context"
	"time"

	"github.com/sullyh7/shahid/internal/store"
	"github.com/sullyh7/shahid/internal/worker"
	"go.uber.org/zap"
)

const (
	path      = "media"
	urlPrefix = "https://www.youtube.com/watch?v="
)

type VideoService struct {
	Store  store.Storage
	FS     FileService
	TS     TranscriptService
	Logger *zap.SugaredLogger
}

func (vs *VideoService) Process(ctx context.Context, j worker.Job) error {
	vs.Logger.Infow("running process", "job", j.VideoID, "url", urlPrefix+j.VideoID)

	time.Sleep(time.Second * 5)
	return nil
}
