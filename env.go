package main

import (
	"os"
	"strconv"
)

func getEnvString(key string, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

func getEnvBool(key string, fallback bool) bool {
	val, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		return fallback
	}
	return val
}

func getEnvInt(key string, fallback int) int {
	val, err := strconv.ParseInt(os.Getenv(key), 10, 32)
	if err != nil {
		return fallback
	}
	return int(val)
}

func getEnvUint(key string, fallback uint) uint {
	val, err := strconv.ParseUint(os.Getenv(key), 10, 32)
	if err != nil {
		return fallback
	}
	return uint(val)
}
