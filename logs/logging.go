package logs

import (
	"os"

	"git.sr.ht/~icikowski/traffic-generator/constants"
	"github.com/rs/zerolog"
)

var console = &filteredWriter{
	w: zerolog.MultiLevelWriter(zerolog.ConsoleWriter{
		NoColor:    false,
		Out:        os.Stdout,
		TimeFormat: constants.LogTimeFormat,
	}),
	l: zerolog.DebugLevel,
}

// SetConsoleWriterLevel specifies the level of logs written to the console
func SetConsoleWriterLevel(l zerolog.Level) {
	console.l = l
}

// Log is global logger instance for logging purposes
var Log = zerolog.New(console).With().Timestamp().Logger()

// LogToFile enables simultaneous logging to both console and given file
func LogToFile(file *os.File) {
	Log = zerolog.New(zerolog.MultiLevelWriter(console, file)).With().Timestamp().Logger()
}
