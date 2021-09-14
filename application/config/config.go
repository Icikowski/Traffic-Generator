package config

import (
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"icikowski.pl/traffic-generator/constants"
	"icikowski.pl/traffic-generator/logs"
)

// Configuration represents traffic generator's parameters
type Configuration struct {
	TrafficName          *string        `json:"name" yaml:"name"`
	Target               *string        `json:"target" yaml:"target"`
	SuccessRatio         *uint          `json:"success_ratio" yaml:"success_ratio"`
	SimultaneousRequests *uint          `json:"simultaneous_requests" yaml:"simultaneous_requests"`
	RequestsInterval     *time.Duration `json:"interval" yaml:"interval"`
	RequestsTimeout      *time.Duration `json:"timeout" yaml:"timeout"`
	InsecureMode         *bool          `json:"insecure" yaml:"insecure"`
}

// Validate checks configuration's correctness and fails if configuration is invalid
func (c *Configuration) Validate() {
	validity := []struct {
		validityCheck func(*Configuration) bool
		message       string
	}{
		{
			validityCheck: func(c *Configuration) bool {
				return c.TrafficName != nil && c.Target != nil &&
					c.SuccessRatio != nil && c.SimultaneousRequests != nil &&
					c.RequestsInterval != nil && c.RequestsTimeout != nil
			},
			message: "configuration has missing fields",
		},
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
			logs.Log.Fatal().Msg(test.message)
		}
	}
	logs.Log.Info().Object("configuration", *c).Msg("configuration is valid")
}

// GetLogFilename returns filename for log file
func (c *Configuration) GetLogFilename() string {
	return fmt.Sprintf(
		"%s (%s).log",
		*c.TrafficName,
		time.Now().Local().Format(constants.FilenameTimeFormat),
	)
}

// Add logging capabilities
var _ zerolog.LogObjectMarshaler = Configuration{}

// MarshalZerologObject adds Configuration's fields to log
func (c Configuration) MarshalZerologObject(e *zerolog.Event) {
	e.
		Str("name", *c.TrafficName).
		Str("target", *c.Target).
		Uint("success_ratio", *c.SuccessRatio).
		Uint("simultaneous_requests", *c.SimultaneousRequests).
		Dur("interval", *c.RequestsInterval).
		Dur("timeout", *c.RequestsTimeout).
		Bool("insecure", *c.InsecureMode)
}
