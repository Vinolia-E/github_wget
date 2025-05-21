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

// Read reads data and reports progress at the specified update rate.
func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.Reader.Read(p)
	pr.BytesRead += int64(n)

	// Only update progress at the specified rate to avoid excessive printing
	if time.Since(pr.LastUpdate) >= pr.UpdateRate {
		elapsedTime := time.Since(pr.StartTime).Seconds()
		speed := float64(0)
		if elapsedTime > 0 {
			speed = float64(pr.BytesRead) / elapsedTime
		}

		var timeRemaining time.Duration
		if speed > 0 && pr.TotalSize > 0 {
			remainingBytes := pr.TotalSize - pr.BytesRead
			remainingSeconds := float64(remainingBytes) / speed
			timeRemaining = time.Duration(remainingSeconds * float64(time.Second))
		}

		pr.ProgressFunc(pr.BytesRead, pr.TotalSize, speed, timeRemaining)
		pr.LastUpdate = time.Now()
	}

	return n, err
}
