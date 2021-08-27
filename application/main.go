package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"icikowski.pl/traffic-generator/logs"
	"icikowski.pl/traffic-generator/types"
	"icikowski.pl/traffic-generator/utils"
)

func main() {
	cycle := 0
	downtime := 0 * time.Second
	totalDowntime := 0 * time.Second

	utils.CreateShutdownHook(func() {
		logs.Log.Info().Msg("shutdown requested by user")
		logs.Log.Info().Str("total downtime", utils.FormatDuration(totalDowntime)).Msg("end statistics")
		logs.Log.Info().Msg("shutdown successful")
	})

	config := types.Configuration{
		TrafficName:          flag.String("name", "some-traffic", "traffic name"),
		Target:               flag.String("target", "", "traffic target"),
		SuccessRatio:         flag.Uint("success", 90, "desired success ratio in percents [1-100]"),
		SimultaneousRequests: flag.Uint("requests", 30, "number of simultaneous requests to be sent in given interval"),
		RequestsInterval:     flag.Duration("interval", 2*time.Second, "requests interval"),
		RequestsTimeout:      flag.Duration("timeout", 1*time.Second, "requests timeout (must not be longer than interval)"),
	}
	verbose := flag.Bool("verbose", false, "enable verbose console logging")
	flag.Parse()

	if !*verbose {
		logs.SetConsoleWriterLevel(zerolog.InfoLevel)
	}

	config.Validate(&logs.Log)

	filename := config.GetLogFilename()
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logs.Log.Error().Str("filename", filename).Err(err).Msg("cannot write log to file")
	} else {
		logs.LogToFile(file)
		logs.Log.Info().Str("filename", filename).Msg("started logging to file")
	}

	logs.Log.Info().Msg("starting analysis")
	for {
		cycle += 1
		logs.Log.Debug().Int("cycle", cycle).Msg("cycle started")

		ctx, cancel := context.WithTimeout(context.Background(), *config.RequestsInterval)
		client := http.Client{
			Timeout: *config.RequestsTimeout,
		}

		mutex := sync.Mutex{}
		successful := 0
		for i := 0; i < int(*config.SimultaneousRequests); i++ {
			go func() {
				req, _ := http.NewRequestWithContext(ctx, http.MethodGet, *config.Target, nil)
				if _, err := client.Do(req); err == nil {
					mutex.Lock()
					defer mutex.Unlock()
					successful += 1
				}
			}()
		}

		<-ctx.Done()
		cancel()

		successRatio := float32(successful*100.0) / float32(*config.SimultaneousRequests)
		if successRatio < float32(*config.SuccessRatio) {
			logs.Log.Debug().Int("cycle", cycle).Uint("ratio requested", *config.SuccessRatio).Float32("ratio actual", successRatio).Msg("success ratio is below requested value")
			if successRatio == 0.0 {
				if downtime == 0 {
					logs.Log.Warn().Int("cycle", cycle).Msg("target is DOWN")
				}
				downtime += *config.RequestsInterval
				totalDowntime += *config.RequestsInterval
			}
		} else if downtime != 0 {
			logs.Log.Info().Int("cycle", cycle).Str("downtime", utils.FormatDuration(downtime)).Msg("target is UP again")
			downtime = 0
		}
		logs.Log.Debug().Int("cycle", cycle).Float32("success ratio", successRatio).Msg("cycle finished")
	}
}
