package lic

import (
	"fmt"

	"github.com/cbroglie/mustache"
)

// Print outputs license metadata to string
func Print(lics []*License, template string) (string, error) {

	var output string
	var err error

	if template != "" {

		// render using custom template file
		output, err = mustache.RenderFile(template, lics)

	} else {

		// use standard text template string
		output, err = mustache.Render(standardTemplate, lics)

	}

	if err != nil {
		return "", fmt.Errorf("template rendering failed: %v", err)
	}

	return output, nil
}

// standard text rendering template
const standardTemplate = `{{#.}}
{{Package}}{{#Version}}@{{Version}}{{/Version}}
{{{License}}}
---
{{/.}}`
