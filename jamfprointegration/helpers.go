package jamfprointegration

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// TODO this duplicated across both integrations. Improve it another time.

// ParseISO8601Date attempts to parse a string date in ISO 8601 format.
func ParseISO8601_Date(dateStr string) (time.Time, error) {
	return time.Parse(time.RFC3339, dateStr)
}

// SafeOpenFile opens a file safely after validating and resolving its path.
func SafeOpenFile(filePath string) (*os.File, error) {
	cleanPath := filepath.Clean(filePath)

	absPath, err := filepath.EvalSymlinks(cleanPath)
	if err != nil {
		return nil, fmt.Errorf("unable to resolve the absolute path: %s, error: %w", filePath, err)
	}

	return os.Open(absPath)
}
