package apis

import (
	"Video_fetcher"
)

//AllVideoData - response struct for all videos api
type AllVideoData struct {
	Success bool                   `json:"success"`
	Data    []*Video_fetcher.Video `json:"data"`
	NextVid string                 `json:"next_vid"`
	Error   string                 `json:"error"`
}

//SearchedVideoData - response struct for search video api
type SearchedVideoData struct {
	Data    *Video_fetcher.Video `json:"data"`
	Success bool                 `json:"success"`
	Error   string               `json:"error"`
}

//Filter - filter struct for user query. More field can be added in this ans support can be extended.
type Filter struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	VideoId     string `json:"video_id"`
}
