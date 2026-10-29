package main

import (
	"os"
	"strconv"
)

func GetEnvInt(def int, key string) int {
	value := os.Getenv(key)

	cv, err := strconv.Atoi(value)

	if err != nil {
		return def
	}

	return cv
}
