package config

import (
	"os"
	"strconv"
	"time"
)

func getString(name string, defaultValue string) string {
	value := os.Getenv(name)

	if value == "" {
		return defaultValue
	}

	return value
}

func getInt(name string, defaultValue int) int {
	stringValue := os.Getenv(name)

	if stringValue == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(stringValue)
	if err != nil {
		panic(err)
	}

	return value
}

func getTimeDuration(name string, defaultValue string) time.Duration {
	stringValue := os.Getenv(name)

	if stringValue == "" {
		return parseDurationString(defaultValue)
	}

	return parseDurationString(defaultValue)
}

func parseDurationString(valueString string) time.Duration {
	value, err := time.ParseDuration(valueString)
	if err != nil {
		panic(err)
	}

	return value
}
