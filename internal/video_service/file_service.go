package videoservice

import (
	"context"
	"io"
)

type FileService interface {
	SaveVideo(context.Context, string, io.Reader) (string, error)
	LoadVideo(context.Context, string) (string, error)
	DeleteVideo(context.Context, string) error

	SaveSubtitle(context.Context, string, []byte) (string, error)
	LoadSubtitle(context.Context, string) (string, error)
	DeleteSubtitle(context.Context, string) error
}

type LocalFileService struct {
}

func (lfs *LocalFileService) SaveVideo(ctx context.Context, videoID string, data io.Reader) (string, error) {
	return "", nil
}
func (lfs *LocalFileService) LoadVideo(ctx context.Context, videoID string) (string, error) {
	return "", nil
}
func (lfs *LocalFileService) DeleteVideo(ctx context.Context, videoID string) error {
	return nil
}
func (lfs *LocalFileService) SaveSubtitle(ctx context.Context, videoID string, data []byte) (string, error) {
	return "", nil
}
func (lfs *LocalFileService) LoadSubtitle(ctx context.Context, videoID string) (string, error) {
	return "", nil
}
func (lfs *LocalFileService) DeleteSubtitle(ctx context.Context, videoID string) error {
	return nil
}
