package apis

import (
	"Video_fetcher"
)

type AllVideoData struct {
	Success bool                   `json:"success"`
	Data    []*Video_fetcher.Video `json:"data"`
	NextVid string                 `json:"next_vid"`
	Error   string                 `json:"error"`
}

type SearchedVideoData struct {
	Data    *Video_fetcher.Video `json:"data"`
	Success bool                 `json:"success"`
	Error   string               `json:"error"`
}

type Filter struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	VideoId     string `json:"video_id"`
}
