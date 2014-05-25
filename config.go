package dogo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

// File that user information such as client id and
// api keys will be written to.
var (
	defaultConfigFile = fmt.Sprintf("%s/.fish", os.ExpandEnv("$HOME"))
	configFile        = ""
	config            Config
)

// Configuration settings.
type Config struct {
	Conf conf
}

// Parsed configuration settings.
type conf struct {
	ClientID   string `toml:"client_id"`
	APIKey     string `toml:"api_key"`
}

// Loads the configuration from the filepath. If there's any
// issue an error will be returned.
func LoadConfig(path string) error {
	if path == "" {
		configFile = defaultConfigFile
	} else {
		configFile = path
	}

	if !fileExists(configFile) {
		return errors.New(fmt.Sprintf("Configuration file not found at %s\n", configFile))
	}

	// The file exists, let's parse it
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	_, err = toml.Decode(string(bytes), &config)
	if err != nil {
		return err
	}

	return nil
}
