package utils

import "regexp"

func MatchStringArray(list []string, str string) bool  {
	for _, item := range list {
		match, _ := regexp.MatchString(item, str)
		if match {
			return true
		}
	}
	return false
}
