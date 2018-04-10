package lic

import "github.com/kogu-io/lisc/proj"

// implicit returns implicit license metadata
func implicit(dep *proj.Project) *License {

	return &License{
		Package:  dep.Name,
		Version:  dep.Version,
		License:  "MIT",
		Explicit: false,
	}

}
