package store

type Video struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Thumbnail   string   `json:"thumbnail"`
	URL         string   `json:"url"`
	Duration    int64    `json:"duration"`
	UploadedAt  string   `json:"uploaded_at"`
	ChannelID   int64    `json:"channel_id"`
	ChannelName string   `json:"channel_name"`
	PlaylistID  int64    `json:"playlist_id,omitempty"`
	Playlist    PlayList `json:"playlist,omitempty"`
}

type PlayList struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	URL         string `json:"url"`
	CreatedAt   string `json:"created_at"`
}
