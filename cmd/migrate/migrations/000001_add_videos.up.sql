CREATE TABLE videos (
    id TEXT NOT NULL PRIMARY KEY,
    channel_id TEXT,
    title TEXT,
    thumbnail_path TEXT,
    video_path TEXT,
    status TEXT NOT NULL CHECK (
        status IN ('pending', 'processing', 'processed', 'failed')
    ),
    error_message TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    processed_at TIMESTAMPTZ
);

CREATE INDEX idx_videos_status ON videos (status);

