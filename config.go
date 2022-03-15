package redis

import (
	"fmt"
)

const (
	// ConfigURI is the location of the redis instance to connect to
	ConfigURI = "uri"
    ConfigChannel = "channel"
)

func requiredConfigErr(name string) error {
	return fmt.Errorf("%q config value must be set", name)
}
