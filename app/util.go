package app

import (
	"fmt"
	"strings"
)

func orElseString[T ~string](value, defaultValue T) T {
	if strings.TrimSpace(string(value)) == "" {
		return defaultValue
	}
	return value
}

func log(format string, args ...any) (int, error) {
	return fmt.Printf(format+"\n", args...)
}
