package logs

import "github.com/rs/zerolog"

type filteredWriter struct {
	w zerolog.LevelWriter
	l zerolog.Level
}

var _ zerolog.LevelWriter = &filteredWriter{}

func (fw *filteredWriter) Write(p []byte) (n int, err error) {
	return fw.w.Write(p)
}

func (fw *filteredWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if fw.l <= level {
		return fw.w.WriteLevel(level, p)
	}
	return len(p), nil
}
