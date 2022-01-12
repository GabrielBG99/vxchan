package rest

import (
	"strconv"
)

// StringToInt tries to convert the input to int or return a default value if input is empty
func StringToInt(i string, def int) (int, error) {
	if i == "" {
		return def, nil
	}

	return strconv.Atoi(i)
}
