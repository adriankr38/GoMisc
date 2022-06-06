package misc

import (
	"strings"
)

func StringInSlice(str string, stringSlice []string, caseSensitive bool) bool {
	if caseSensitive {
		for i := 0; i < len(stringSlice); i++ {
			if str == stringSlice[i] {
				return true
			}

		}

		return false
	}

	for i := 0; i < len(stringSlice); i++ {
		if strings.ToLower(str) == strings.ToLower(stringSlice[i]) {
			return true
		}

	}

	return false
}

