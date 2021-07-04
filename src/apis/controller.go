package apis

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"utils"
)

// GetVideos - this controller will return all cached videos in paginated form. Current page size is 10.
func GetVideos(ginContext *gin.Context) {
	serializer := utils.GetSerializer("json")
	defer func() {
		if r := recover(); r != nil {
			serializer.ReturnError(ginContext, http.StatusInternalServerError, "Error: %v", r)
		}
	}()
	nextVId := ginContext.DefaultQuery("next", "")
	resultData := getAllVideos(nextVId)
	if resultByte, marshallErr := json.Marshal(resultData); marshallErr == nil {
		serializer.ReturnOk(ginContext, resultByte)
	} else {
		serializer.ReturnError(ginContext, http.StatusInternalServerError, "Error: %v", marshallErr)
	}
}

// SearchVideo - this controller will search cached video's data with user searched query string and return response.
func SearchVideo(ginContext *gin.Context) {
	serializer := utils.GetSerializer("json")
	defer func() {
		if r := recover(); r != nil {
			serializer.ReturnError(ginContext, http.StatusInternalServerError, "Error: %v", r)
		}
	}()
	//extract filter dict
	filterString := ginContext.Query("f")
	filter := new(Filter)
	if filterString != "" {
		if err := json.Unmarshal([]byte(filterString), filter); err != nil {
			serializer.ReturnError(ginContext, http.StatusBadRequest, "Error: %v", err)
		}
	}
	resultData := getSearchedVideo(filter)
	if resultByte, marshallErr := json.Marshal(resultData); marshallErr == nil {
		serializer.ReturnOk(ginContext, resultByte)
	} else {
		serializer.ReturnError(ginContext, http.StatusInternalServerError, "Error: %v", marshallErr)
	}
}
