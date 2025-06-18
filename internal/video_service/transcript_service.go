package videoservice

import "context"

type TranscriptResult struct {
	SRT      []byte
	Text     string
	Segments []TranscriptSegment
}

type TranscriptSegment struct {
	Start float64
	End   float64
	Text  string
}

type TranscriptService interface {
	Transcribe(context.Context, string) (*TranscriptResult, error)
}

type StandardTranscriptService struct {
}

func (std *StandardTranscriptService) Transcribe(ctx context.Context, videoPath string) (*TranscriptResult, error) {
	return nil, nil
}
