package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"constants"
)

type Serializer struct {
	contentType string
}

var (
	jsonSerializer = &Serializer{
		contentType: constants.ACCEPT_HEADERS_JSON,
	}
	protoSerializer = &Serializer{
		contentType: constants.ACCEPT_HEADERS_PROTO,
	}
)

func GetSerializer(serializerType string) *Serializer {
	switch serializerType {
	case "proto":
		return protoSerializer
	default:
		return jsonSerializer
	}
}

func IsShowErrors(ginContext *gin.Context) bool {
	errorHash := ginContext.Query("errorHash")
	if errorHash == constants.ERROR_HASH {
		return true
	}
	return false
}

func (serializer *Serializer) ReturnOk(ginContext *gin.Context, result []byte) {
	if ginContext.IsAborted() {
		return
	}
	ginContext.Abort()
	showErrors := IsShowErrors(ginContext)
	if !showErrors {
		ginContext.Data(http.StatusOK, serializer.contentType, result)
	} else {
		errorsJson, _ := json.Marshal(ginContext.Errors)
		ginContext.Data(http.StatusOK, serializer.contentType, errorsJson)
	}
}

func (serializer *Serializer) ReturnError(ginContext *gin.Context, code int, sfmt string, v ...interface{}) {
	if ginContext.IsAborted() {
		return
	}
	ginContext.Abort()
	showErrors := IsShowErrors(ginContext)
	if !showErrors {
		ginContext.Data(code, serializer.contentType, []byte(fmt.Sprintf("{\"error\":\"%s\"}", fmt.Sprintf(sfmt, v...))))
	} else {
		errorsJson, _ := json.Marshal(ginContext.Errors)
		ginContext.Data(code, serializer.contentType, errorsJson)
	}
}

func (serializer *Serializer) ReturnErrorCode(ginContext *gin.Context, code int, result []byte) {
	if ginContext.IsAborted() {
		return
	}
	ginContext.Abort()
	showErrors := IsShowErrors(ginContext)
	if !showErrors {
		ginContext.Data(code, serializer.contentType, result)
	} else {
		errorsJson, _ := json.Marshal(ginContext.Errors)
		ginContext.Data(code, serializer.contentType, errorsJson)
	}
}
