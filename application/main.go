package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"flag"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"icikowski.pl/traffic-generator/config"
	"icikowski.pl/traffic-generator/constants"
	"icikowski.pl/traffic-generator/logs"
	"icikowski.pl/traffic-generator/utils"
)

var version = constants.BuildValueUnknown
var gitCommit = constants.BuildValueUnknown
var binaryType = constants.BuildValueUnknown

func main() {
	cycle := 0
	downtime := 0 * time.Second
	totalDowntime := 0 * time.Second

	utils.RegisterGracefulExitHooks(func() {
		logs.Log.Info().Msg("user requested application exit")
		logs.Log.Info().Str("total downtime", utils.FormatDuration(totalDowntime)).Msg("prepared statistics")
		logs.Log.Info().Msg("application finished")
	})

	directConfig := &config.Configuration{
		TrafficName:          flag.String("name", "some-traffic", "traffic name"),
		Target:               flag.String("target", "", "traffic target"),
		SuccessRatio:         flag.Uint("success", 90, "desired success ratio in percents [1-100]"),
		SimultaneousRequests: flag.Uint("requests", 30, "number of simultaneous requests to be sent in given interval"),
		RequestsInterval:     flag.Duration("interval", 2*time.Second, "requests interval"),
		RequestsTimeout:      flag.Duration("timeout", 1*time.Second, "requests timeout (must not be longer than interval)"),
		InsecureMode:         flag.Bool("insecure", false, "insecure mode (SSL certificates of the target will not be verified)"),
	}
	configInput := flag.String("config", "", "configuration file (YAML or JSON) or pipeline input (\"--\")")
	verbose := flag.Bool("verbose", false, "enable verbose console logging")
	versionCmd := flag.Bool("version", false, "print application's version and build info")
	flag.Parse()

	if *versionCmd {
		out, _ := json.Marshal(map[string]string{
			"version":    version,
			"commit":     gitCommit,
			"binaryType": binaryType,
		})
		os.Stdout.Write(out)
		return
	}

	if !*verbose {
		logs.SetConsoleWriterLevel(zerolog.InfoLevel)
	}

	var runningConfig *config.Configuration
	runningConfig, err := config.LoadExternalConfig(*configInput)
	if err != nil {
		if err != config.ErrEmptyFlag {
			logs.Log.Fatal().Err(err).Msg("provided config is invalid")
		}
		logs.Log.Warn().Err(err).Msg("")
		runningConfig = directConfig
	}

	runningConfig.Validate()

	filename := runningConfig.GetLogFilename()
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

		ctx, cancel := context.WithTimeout(context.Background(), *runningConfig.RequestsInterval)
		client := http.Client{
			Timeout: *runningConfig.RequestsTimeout,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: *runningConfig.InsecureMode,
				},
			},
		}

		mutex := sync.Mutex{}
		successful := 0
		for i := 0; i < int(*runningConfig.SimultaneousRequests); i++ {
			go func() {
				req, _ := http.NewRequestWithContext(ctx, http.MethodGet, *runningConfig.Target, nil)
				if _, err := client.Do(req); err == nil {
					mutex.Lock()
					defer mutex.Unlock()
					successful += 1
				}
			}()
		}

		<-ctx.Done()
		cancel()

		successRatio := float32(successful*100.0) / float32(*runningConfig.SimultaneousRequests)
		if successRatio < float32(*runningConfig.SuccessRatio) {
			logs.Log.Debug().Int("cycle", cycle).Uint("ratio requested", *runningConfig.SuccessRatio).Float32("ratio actual", successRatio).Msg("success ratio is below requested value")
			if successRatio == 0.0 {
				if downtime == 0 {
					logs.Log.Warn().Int("cycle", cycle).Msg("target is DOWN")
				}
				downtime += *runningConfig.RequestsInterval
				totalDowntime += *runningConfig.RequestsInterval
			}
		} else if downtime != 0 {
			logs.Log.Info().Int("cycle", cycle).Str("downtime", utils.FormatDuration(downtime)).Msg("target is UP again")
			downtime = 0
		}
		logs.Log.Debug().Int("cycle", cycle).Float32("success ratio", successRatio).Msg("cycle finished")
	}
}
