package visualizer

import "errors"

// Common errors
var (
	ErrDataStructureNotFound = errors.New("data structure not found")
	ErrInvalidOperation      = errors.New("invalid operation")
	ErrSnapshotFailed        = errors.New("failed to create snapshot")
)