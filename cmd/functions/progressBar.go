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
