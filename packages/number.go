package packages

import (
	"strconv"
)

func ParseIntMin(def, min int, val string) int {
	cv, err := strconv.Atoi(val)

	if err != nil {
		return def
	}

	if cv <= min {
		return min
	}

	return cv
}
