package app

import (
	"fmt"
	"strings"
)

func stringOrDefault[T ~string](value, defaultValue T) T {
	if strings.TrimSpace(string(value)) == "" {
		return defaultValue
	}
	return value
}

func joinLines(lines ...string) string {
	return strings.Join(lines, "\n")
}

func log(format string, args ...any) (int, error) {
	return fmt.Printf(format+"\n", args...)
}
