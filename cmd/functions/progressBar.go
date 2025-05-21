package functions

import (
	"io"
	"time"
)

type ProgressReader struct {
	Reader       io.Reader
	TotalSize    int64
	BytesRead    int64
	StartTime    time.Time
	LastUpdate   time.Time
	UpdateRate   time.Duration
	ProgressFunc func(bytesRead, totalSize int64, speed float64, timeRemaining time.Duration)
}

// ProgressReader wraps an io.Reader to track and report progress over time.
func NewProgressReader(reader io.Reader, totalSize int64, progressFunc func(bytesRead, totalSize int64, speed float64, timeRemaining time.Duration)) *ProgressReader {
	return &ProgressReader{
		Reader:       reader,
		TotalSize:    totalSize,
		BytesRead:    0,
		StartTime:    time.Now(),
		LastUpdate:   time.Now(),
		UpdateRate:   time.Millisecond * 100, // Update progress every 100ms
		ProgressFunc: progressFunc,
	}
}
