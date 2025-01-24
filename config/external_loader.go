package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"git.sr.ht/~icikowski/traffic-generator/logs"
	"gopkg.in/yaml.v3"
)

var (
	// ErrEmptyFlag is an error returned when the config file is not specified by corresponding flag
	ErrEmptyFlag = errors.New("configuration flag has empty value, falling back to values passed directly")
	// ErrNotReadable is an error returned when the config file or direct input cannot be read
	ErrNotReadable = errors.New("configuration is not readable")
	// ErrNotParseable is an error returned when the config file or direct input cannot be parsed
	ErrNotParseable = errors.New("configuration cannot be parsed as neither YAML nor JSON")
)

func LoadExternalConfig(filename string) (*Configuration, error) {
	if filename == "" {
		return nil, ErrEmptyFlag
	}
	logs.Log.Debug().Str("filename", filename).Msg("provided config input")

	var contents []byte
	var err error
	if filename == "--" {
		contents, err = ioutil.ReadAll(os.Stdin)
	} else {
		contents, err = ioutil.ReadFile(filename)
	}

	if err != nil {
		logs.Log.Debug().Err(err).Msg("configuration loader met unexpected error")
		return nil, ErrNotReadable
	}

	parsedConfig := &Configuration{}
	switch {
	case yaml.Unmarshal(contents, parsedConfig) == nil:
		logs.Log.Debug().Msg("configuration loader loaded YAML input")
	case json.Unmarshal(contents, parsedConfig) == nil:
		logs.Log.Debug().Msg("configuration loader loaded JSON input")
	default:
		logs.Log.Debug().Msg("configuration loader couldn't parse load as YAML/JSON")
		return nil, ErrNotParseable
	}

	return parsedConfig, nil
}
