package utils

import (
	"log"
)

func HandleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}

// IsValueInList - search for a value in list. Return true/false
func IsValueInList(value string, list []string) (bool, int) {
	for index, v := range list {
		if v == value {
			return true, index
		}
	}
	return false, 0
}
