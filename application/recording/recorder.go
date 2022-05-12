package recording

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
)

// Recorder is an object that is able to record success ratio
// in separate goroutine
type Recorder struct {
	queue  chan []string
	writer *csv.Writer
	ctx    context.Context
	cancel context.CancelFunc
}

// Record records the success ratio of given time
func (r *Recorder) Record(t, successRatio float64) {
	go func(t, successRatio float64) {
		timeStr := fmt.Sprintf("%f", t)
		successRatioStr := fmt.Sprintf("%f", successRatio)
		r.queue <- []string{timeStr, successRatioStr}
	}(t, successRatio)
}

func (r *Recorder) Stop() {
	r.cancel()
}

func (r *Recorder) maintain() {
	for {
		select {
		case <-r.ctx.Done():
			r.writer.Flush()
			return
		case entry := <-r.queue:
			r.writer.Write(entry)
			r.writer.Flush()
		}
	}
}

// NewRecorder creates new recorder
func NewRecorder(filename string) (*Recorder, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	writer := csv.NewWriter(file)
	ctx, cancel := context.WithCancel(context.Background())

	recorder := &Recorder{
		queue:  make(chan []string),
		writer: writer,
		ctx:    ctx,
		cancel: cancel,
	}

	go recorder.maintain()
	go func() {
		recorder.queue <- []string{"time", "success ratio"}
	}()

	return recorder, nil
}
