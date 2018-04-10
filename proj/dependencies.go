package proj

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Dependencies extracts go project's dependencies metadata
func Dependencies(dir string) (*List, error) {

	// project's Gopkg.lock location
	path := filepath.Join(dir, "Gopkg.lock")

	// open Gopkg.lock for reading
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to access file: %s", path)
	}

	// stream parse TOML content
	deps := &List{}
	_, err = toml.DecodeReader(f, deps)
	if err != nil {
		return nil, fmt.Errorf("unable to parse file content as TOML: %s", path)
	}

	return deps, nil
}
