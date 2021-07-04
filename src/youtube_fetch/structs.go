package youtube_fetch

type Youtube struct {
	VideoId      string    `json:"video_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Url          string    `json:"url"`
	ChannelTitle string    `json:"channel_title"`
	ChannelId    string    `json:"channel_id"`
	PublishedAt  string `json:"published_at"`
}

func (yt *Youtube)GetVideoUrl() (url string) {
	if yt!=nil {
		return "https://www.youtube.com/watch?v=" + yt.VideoId
	}
	return
}