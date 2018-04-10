package cmd

import (
	"fmt"
	"os"

	"github.com/kogu-io/lisc/lic"
	"github.com/kogu-io/lisc/proj"
	"github.com/spf13/cobra"
)

// command
const cmd = "lisc"

// descriptions
const descShort = "lisc scans and combines vendor dependencies licenses in dep-aware projects"
const descLong = `lisc scans and combines vendor dependencies licenses in dep-aware projects.
It fetches packages list from 'Gopkg.lock' and searches licenses in the corresponding 'vendor/' subfolders.
The retrieved metadata is processed by a mustache template, and printed to standard output.`

// parameter values
var root string
var template string
var lookup []string
var ensure bool

// errors
const errorDependenciesDetecting = 1
const errorLicenseScanning = 2
const errorLicenseRendering = 3

func init() {

	// configure flags
	rootCmd.Flags().StringVarP(&root, "source", "s", ".", "Source directory to read Gopkg.lock from")
	rootCmd.Flags().StringVarP(&template, "template", "t", "", "Path to the mustache template file to use for output generation")
	rootCmd.Flags().StringSliceVarP(&lookup, "names", "n", []string{}, "Extra lookup file names to extract license from")
	rootCmd.Flags().BoolVarP(&ensure, "ensure", "e", false, "Ensure all dependencies' licenses are detected")
}

var rootCmd = &cobra.Command{
	Use: cmd, Short: descShort, Long: descLong,
	Run: func(cmd *cobra.Command, args []string) {

		// detect dependencies
		deps, err := proj.Dependencies(root)
		if err != nil {
			fmt.Printf("Dependencies detection failed: %v\n", err)
			os.Exit(errorDependenciesDetecting)
		}

		// fetch licenses
		lics, err := lic.Scan(root, deps, lookup, ensure)
		if err != nil {
			fmt.Printf("License scan failed: %v\n", err)
			os.Exit(errorLicenseScanning)
		}

		// combine licenses
		text, err := lic.Print(lics, template)
		if err != nil {
			fmt.Printf("License printing failed: %v\n", err)
			os.Exit(errorLicenseRendering)
		}

		// output to console
		fmt.Print(text)
	},
}

// Execute starts CLI handling
func Execute() error {
	return rootCmd.Execute()
}
