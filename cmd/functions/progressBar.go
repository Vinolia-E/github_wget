package functions

import (
	"fmt"
	"io"
	"strings"
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

// DisplayProgress renders a textual progress bar and download metrics.
func DisplayProgress(bytesRead, totalSize int64, speed float64, timeRemaining time.Duration) {
	// Calculate percentage
	percentage := float64(100)
	if totalSize > 0 {
		percentage = float64(bytesRead) * 100 / float64(totalSize)
	}

	// Format sizes
	var readSize, totalSizeStr string
	if bytesRead < 1024*1024 {
		readSize = fmt.Sprintf("%.2f KiB", float64(bytesRead)/1024)
	} else {
		readSize = fmt.Sprintf("%.2f MiB", float64(bytesRead)/(1024*1024))
	}

	if totalSize < 1024*1024 {
		totalSizeStr = fmt.Sprintf("%.2f KiB", float64(totalSize)/1024)
	} else {
		totalSizeStr = fmt.Sprintf("%.2f MiB", float64(totalSize)/(1024*1024))
	}

	// Format speed
	var speedStr string
	if speed < 1024*1024 {
		speedStr = fmt.Sprintf("%.2f KiB/s", speed/1024)
	} else {
		speedStr = fmt.Sprintf("%.2f MiB/s", speed/(1024*1024))
	}

	// Format time remaining
	var timeStr string
	if timeRemaining > time.Hour {
		timeStr = fmt.Sprintf("%dh%dm", int(timeRemaining.Hours()), int(timeRemaining.Minutes())%60)
	} else if timeRemaining > time.Minute {
		timeStr = fmt.Sprintf("%dm%ds", int(timeRemaining.Minutes()), int(timeRemaining.Seconds())%60)
	} else {
		timeStr = fmt.Sprintf("%ds", int(timeRemaining.Seconds()))
	}

	// Create progress bar (50 characters wide)
	width := 50
	completed := int(float64(width) * percentage / 100)
	bar := strings.Repeat("=", completed) + strings.Repeat(" ", width-completed)

	// Clear the current line and print the progress
	fmt.Printf("\r %s / %s [%s] %.2f%% %s %s",
		readSize, totalSizeStr, bar, percentage, speedStr, timeStr)

	// If download is complete, add a newline
	if bytesRead >= totalSize {
		fmt.Println()
	}
}

// DownloadWithProgress copies data from reader to writer while displaying progress.
func DownloadWithProgress(reader io.Reader, writer io.Writer, totalSize int64) (int64, error) {
	progressReader := NewProgressReader(reader, totalSize, DisplayProgress)
	return io.Copy(writer, progressReader)
}
