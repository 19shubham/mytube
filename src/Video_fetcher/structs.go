package Video_fetcher

//Video - a struct for storing and using fields related to particular video. In this we have assumed videoId will be
// unique for each video.
type Video struct {
	VideoId      string `json:"video_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Url          string `json:"url"`
	ChannelTitle string `json:"channel_title"`
	ChannelId    string `json:"channel_id"`
	PublishedAt  string `json:"published_at"`
}

func (yt *Video) GetVideoUrl() (url string) {
	if yt != nil {
		return "https://www.youtube.com/watch?v=" + yt.VideoId
	}
	return
}
