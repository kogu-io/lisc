package lic

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/kogu-io/lisc/proj"
)

// prioritized list of file names
// to search for license content
// (generated from names + extensions)
var usualLookupNames []string

var usualLicenseFileNames = [...]string{
	"license", "licence", "copying",
}

var usualLicenseFileExtensions = [...]string{
	"", ".md", ".txt",
}

func init() {

	// combine usual license file names lookup list
	for _, name := range usualLicenseFileNames {
		for _, ext := range usualLicenseFileExtensions {
			usualLookupNames = append(usualLookupNames, name+ext)
		}
	}

}

// Scan fetches project dependencies' license metadata
func Scan(root string, deps *proj.List, extraLookup []string, ensure bool) ([]*License, error) {

	// combine lookup names
	lookup := lookupNames(extraLookup)

	licenses := []*License{}

	// iterate over all dependencies
	for _, dep := range deps.Items {

		lic, err := license(root, dep, lookup)
		if err != nil {
			return nil, err
		}

		// validate license is found if strict mode requested
		if ensure && !lic.Explicit {
			return nil, fmt.Errorf("license not found: %s", dep.Name)
		}

		licenses = append(licenses, lic)
	}

	return licenses, nil
}

func lookupNames(extraLookup []string) []string {

	// copy default lookup name list
	lookup := make([]string, len(usualLookupNames))
	copy(lookup, usualLookupNames)

	// append extra lookup names to the bottom
	lookup = append(lookup, extraLookup...)
	return lookup
}

func vendoredPath(root string, dep *proj.Project) string {
	return filepath.Join(root, "vendor", dep.Name)
}

func license(root string, dep *proj.Project, lookup []string) (*License, error) {

	// default to implicit license
	lic := implicit(dep)

	// construct dependency project path in /vendor/ folder
	path := vendoredPath(root, dep)

	// fetch license text
	text, err := licenseText(path, lookup)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve license: %v", err)
	}

	// apply license if found
	if text != "" {
		lic.License = text
		lic.Explicit = true
	}

	return lic, nil
}

func licenseText(dir string, lookup []string) (string, error) {

	// list all files inside project root directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", fmt.Errorf("unable to access dependent project directory: %s", dir)
	}

	// sequentially search for expected license files
	for _, name := range lookup {
		for _, file := range files {

			// case-insensitive name comparison
			if strings.EqualFold(name, file.Name()) {

				path := filepath.Join(dir, file.Name())

				// read the detected license file content
				bytes, err := ioutil.ReadFile(path)
				if err == nil && len(bytes) > 0 {
					return string(bytes), nil
				}

				// in case read operation failure,
				// continue searching for alternative files
			}
		}
	}

	return "", nil
}
