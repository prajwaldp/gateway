package lib

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

// Config is a struct that holds the configuration of the gateway.
type Config struct {
	// Port that the gateway should listen on.
	Port string `json:"port"`

	//  Versions maps from the version (eg. 'v1') to path prefix to backend service, e.g. /users -> users-api.mycompany.net. The full path will be
	//  passed through to the backend service, i.e. in the above example users-api.mycompany.net will see /users as
	//  the first part of the URL.
	Versions map[string]map[string]string `json:"versions"`

	//  The response to return when an unmapped route is requested - an arbitrary JSON object that will be marshalled.
	NotFoundResponse interface{} `json:"not_found_error"`

	//  Fallback URL if the version string is not used - this is commonly the documentation page for the APIs.
	//  The fallback URL will only be used if
	//  - The `version` configuration is not blank
	//  - The URL does not start with the version
	//  - The fallback rule is not blank
	FallbackRule string `json:"fallback_rule"`

	//  Scheme is the type of URL scheme to use for requests to the backend, such as "http" or "https".
	Scheme string `json:"scheme"`
}

// Load loads a configuration file and parses it into a Config struct.
func Load(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return load(b)
}

func load(b []byte) (*Config, error) {
	var config Config
	err := json.Unmarshal(b, &config)
	if err != nil {
		return nil, err
	}
	for k := range config.Versions {
		config.Versions[strings.ToLower(k)] = config.Versions[k] //make version string case insensitive
	}
	return &config, nil
}
