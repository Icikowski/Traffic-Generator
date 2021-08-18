package types

import (
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"icikowski.pl/traffic-generator/constants"
)

// Configuration represents traffic generator's parameters
type Configuration struct {
	TrafficName          *string
	Target               *string
	SuccessRatio         *uint
	SimultaneousRequests *uint
	RequestsInterval     *time.Duration
	RequestsTimeout      *time.Duration
}

// Validate checks configuration's correctness and fails if configuration is invalid
func (c *Configuration) Validate(log *zerolog.Logger) {
	validity := []struct {
		validityCheck func(*Configuration) bool
		message       string
	}{
		{
			validityCheck: func(c *Configuration) bool {
				return len(*c.Target) != 0
			},
			message: "traffic destination not set",
		},
		{
			validityCheck: func(c *Configuration) bool {
				return *c.RequestsTimeout <= *c.RequestsInterval
			},
			message: "requests timeout must not be longer than requests interval",
		},
		{
			validityCheck: func(c *Configuration) bool {
				return *c.SuccessRatio >= 1 && *c.SuccessRatio <= 100
			},
			message: "success ratio should be between [1-100]",
		},
	}

	for _, test := range validity {
		if !test.validityCheck(c) {
			log.Fatal().Msg(test.message)
		}
	}
	log.Info().Dict(
		"config",
		zerolog.Dict().
			Str("traffic name", *c.TrafficName).
			Str("destination", *c.Target).
			Uint("success ratio", *c.SuccessRatio).
			Uint("simultaneous requests", *c.SimultaneousRequests).
			Dur("requests interval", *c.RequestsInterval).
			Dur("requests timeout", *c.RequestsTimeout),
	).Msg("configuration is valid")
}

// GetLogFilename returns filename for log file
func (c *Configuration) GetLogFilename() string {
	return fmt.Sprintf(
		"%s (%s).log",
		*c.TrafficName,
		time.Now().Local().Format(constants.FilenameTimeFormat),
	)
}
