package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/cbroglie/mustache"
	"github.com/spf13/cobra"
)

var source string
var template string

var rootCmd = &cobra.Command{
	Use:   "lisc",
	Short: "lisc outputs licenses from vendor-packages in projects using dep",
	Long: `lisc is naiive implementation for scanning and outputting licenses 
			from used vendor-packages in golang projects that use 
			dep (https://golang.github.io/dep/) for dependency management.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := doStuff(); err != nil {
			fmt.Println(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

const fname = "Gopkg.lock"
const templ = "template.mustache"

func init() {
	rootCmd.Flags().StringVarP(&source, "source", "s", ".", "Source directory to read Gopkg.lock from.")
	rootCmd.Flags().StringVarP(&template, "template", "t", templ, "Path pointing to the mustache template file to use for output generation.")
}

type project struct {
	Name     string
	Packages []string
	Revision string
	Version  string
}

type projects struct {
	Project []project `toml:"projects"`
}

type license struct {
	Package  string
	Version  string
	License  string
	Explicit bool
}

func getProjects(fname string) (*projects, error) {
	var projects projects

	content, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, fmt.Errorf("Cant read %v", fname)
	}

	if _, err := toml.Decode(string(content), &projects); err != nil {
		return nil, fmt.Errorf("Failed to decode %v", fname)
	}

	return &projects, nil
}

var files = [...]string{"LICENSE", "COPYING", "LICENCE"}
var extensions = [...]string{"", ".md", ".txt"}

func getLicenses(c projects, root string) ([]license, error) {
	var licenses []license

	for _, proj := range c.Project {

		lic := license{Package: proj.Name, Version: proj.Version, License: "MIT", Explicit: false}

		for _, fn := range files {
			for _, ex := range extensions {

				fname := fmt.Sprintf("%s%s", fn, ex)
				p := filepath.Join(root, "vendor", proj.Name, fname)

				k, err := ioutil.ReadFile(p)
				if err != nil {
					continue
				}
				lic.License = string(k)
				lic.Explicit = true
				break
			}
		}
		licenses = append(licenses, lic)
	}

	return licenses, nil
}

func doStuff() error {

	passed := filepath.Join(source, fname)
	c, err := getProjects(passed)
	if err != nil {
		return err
	}

	root, _ := filepath.Abs(filepath.Dir(passed))

	licenses, err := getLicenses(*c, root)
	if err != nil {
		return fmt.Errorf("Failed to parse licenses from %v", root)
	}

	content, err := mustache.RenderFile(template, licenses)
	if err != nil {
		return fmt.Errorf("Failed to render %v", template)
	}

	fmt.Printf("%v", content)

	return nil

}
