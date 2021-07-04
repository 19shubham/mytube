package apis

import (
	"sort"

	"Video_fetcher"
	"constants"
	"utils"
)

//getAllVideos - function will format response for all videos api.
func getAllVideos(nextVId string) (result *AllVideoData) {
	result = new(AllVideoData)
	if len(Video_fetcher.StoredData) == 0 {
		result.Error = "Stored Data is Empty."
	}
	if nextVId == constants.END_NEXT_VID {
		result.Error = "No more stored video"
	}
	if result.Error != "" {
		return
	}
	return getVideosFromNextVId(nextVId)
}

//getVideosFromNextVId - helper function for videos api to extract nextvid and do processing and combine response.
func getVideosFromNextVId(nextVId string) (result *AllVideoData) {
	result = new(AllVideoData)
	var (
		vidList    []string
		resNextVid string
	)
	if nextVId == "" {
		for index, vId := range Video_fetcher.OrderedVIds {
			if index < constants.PAGESIZE {
				vidList = append(vidList, vId)
			} else {
				resNextVid = vId
				break
			}
		}
	} else {
		if success, vIDIndex := utils.IsValueInList(nextVId, Video_fetcher.OrderedVIds); success {
			var remainingVIdList = make([]string, 0, len(Video_fetcher.OrderedVIds))
			remainingVIdList = Video_fetcher.OrderedVIds[vIDIndex:]
			for index, vId := range remainingVIdList {
				if index < constants.PAGESIZE {
					vidList = append(vidList, vId)
				} else {
					resNextVid = vId
					break
				}
			}
		} else {
			result.Error = constants.INVALID_QUERY
			return
		}
	}
	for _, videoId := range vidList {
		if data, exists := Video_fetcher.StoredData[videoId]; exists {
			result.Data = append(result.Data, data)
		}
	}
	var list PairList
	list = result.Data
	sort.Sort(list)
	if resNextVid == "" {
		resNextVid = constants.END_NEXT_VID
	}
	if len(result.Data) > 0 {
		result.Success = true
		result.NextVid = resNextVid
		result.Data = list
	}
	return
}

//getSearchedVideo - function will search video on the basis of user searched query and return if found.
func getSearchedVideo(filter *Filter) (result *SearchedVideoData) {
	result = new(SearchedVideoData)
	if filter == nil || (filter.Title == "" && filter.Description == "" && filter.VideoId == "") {
		result.Error = constants.INVALID_QUERY
		return
	}
	for _, videoId := range Video_fetcher.OrderedVIds {
		if data, exists := Video_fetcher.StoredData[videoId]; exists {
			if filter.Title != "" && data.Title == filter.Title {
				result.Data = data
				result.Success = true
				return
			}
			if filter.Description != "" && data.Description == filter.Description {
				result.Data = data
				result.Success = true
				return
			}
			if filter.VideoId != "" && data.VideoId == filter.VideoId {
				result.Data = data
				result.Success = true
				return
			}
		}
	}
	if result.Data == nil {
		result.Error = "Not found in stored data"
	}
	return result
}

//For sorting video list
type PairList []*Video_fetcher.Video

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].PublishedAt > p[j].PublishedAt }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
