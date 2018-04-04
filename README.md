# lisc - LIcense SCanner

lisc is naiive implementation for scanning and outputting licenses from used vendor-packages in golang projects that use dep (https://golang.github.io/dep/) for dependency management. It scans packages located in vendor subfolder that are referenced in Gopkg.lock file and writes the content from LICENSE files to stdout.

Usage:
```
lisc --help

```
