package utils

import (
	"log"
)

//HandleError - utility function to handle error
func HandleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}

// IsValueInList - utility function to search for a value in list. Return true/false and index if value found otherwise 0.
func IsValueInList(value string, list []string) (bool, int) {
	for index, v := range list {
		if v == value {
			return true, index
		}
	}
	return false, 0
}
