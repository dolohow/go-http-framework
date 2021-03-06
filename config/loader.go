// Package config provides the ability to load yaml based config file
// depending on environment variable "GO_ENV".
package config

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	ErrMissingKey    = errors.New("No config filename for the given GO_ENV")
	ErrNoSuchFile    = errors.New("Could not load config file")
	ErrCouldNotParse = errors.New("Could not parse config file")
)

// Filenames type is used to abstract underlying data type. The key
// represents current value of GO_ENV environment variable and a value
// filename in yaml format.
type Filenames map[string]string

type config struct {
	env string
}

// NewConfigLoader returns a new config.
func NewConfigLoader() *config {
	return &config{env: getEnvVariable()}
}

// Load function unmarshales file from files map to v interface.
func (c *config) Load(v interface{}, files Filenames) error {
	fileName, ok := files[c.env]

	if !ok {
		return ErrMissingKey
	}

	config, err := ioutil.ReadFile(fileName)

	if err != nil {
		return ErrNoSuchFile
	}

	if err = yaml.Unmarshal(config, v); err != nil {
		return ErrCouldNotParse
	}

	return nil
}

func getEnvVariable() string {
	if env := os.Getenv("GO_ENV"); env != "" {
		return env
	}

	return "development"
}
