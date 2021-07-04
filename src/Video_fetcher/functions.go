package Video_fetcher

import (
	"time"

	"google.golang.org/api/youtube/v3"

	"constants"
	"utils"
)

var (
	YtService   *youtube.Service  //global YtService variable
	StoredData  map[string]*Video //global stored data map with key as video id
	OrderedVIds []string          //sorted list of video id's by their published time in descending order
)

//FetchVideoData - main video fetcher which will call youtube api and return response with predefined search query
//and order by date. Query can be changed in constants.QUERY_STRING. Currently it is being set to "cricket".
//Max response size requested is 50 in one call, which also can be changed in constants.MAX_RESULT_SIZE.
func FetchVideoData(service *youtube.Service, part []string, isFirstCall bool) {
	if len(StoredData) == 0 {
		StoredData = make(map[string]*Video)
	}
	var then time.Time
	now := time.Now()
	if isFirstCall {
		then = now.Add(time.Duration(-60) * time.Hour)
	} else {
		then = now.Add(time.Duration(-30) * time.Minute)
	}
	publishedAfter := then.Format("2006-01-02T15:04:05Z")
	call := service.Search.List(part, constants.QUERY_STRING, constants.ORDER_BY, constants.MAX_RESULT_SIZE, publishedAfter)
	response, err := call.Do()
	utils.HandleError(err, "")
	if response != nil {
		for _, item := range response.Items {
			if item != nil && item.Id != nil && item.Snippet != nil {
				record := &Video{
					VideoId:      item.Id.VideoId,
					Title:        item.Snippet.Title,
					Description:  item.Snippet.Description,
					Url:          item.Snippet.Thumbnails.Default.Url,
					ChannelTitle: item.Snippet.ChannelTitle,
					ChannelId:    item.Snippet.ChannelId,
					PublishedAt:  item.Snippet.PublishedAt,
				}
				StoredData[item.Id.VideoId] = record
				if isFirstCall {
					OrderedVIds = append(OrderedVIds, item.Id.VideoId)
				} else {
					OrderedVIds = append([]string{item.Id.VideoId}, OrderedVIds...)
				}
			}
		}
	}
	return
}
