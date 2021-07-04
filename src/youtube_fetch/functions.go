package youtube_fetch

import (
	"log"

	"google.golang.org/api/youtube/v3"
)

func FetchSearchData(service *youtube.Service, part []string) (data []*Youtube){
	data= make([]*Youtube,0)
	call := service.Search.List(part)
	response, err := call.Do()
	handleError(err, "")
	for _,item := range response.Items{
		data = append(data, &Youtube{
			VideoId:      item.Id.VideoId,
			Title:        item.Snippet.Title,
			Description:  item.Snippet.Description,
			Url:          item.Snippet.Thumbnails.Default.Url,
			ChannelTitle: item.Snippet.ChannelTitle,
			ChannelId:    item.Snippet.ChannelId,
			PublishedAt:  item.Snippet.PublishedAt,
		})
	}
	return
}

func handleError(err error, message string) {
	if message == "" {message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message + ": %v", err.Error())
	}
}