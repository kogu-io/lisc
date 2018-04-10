# lisc - LIcense SCanner
lisc scans and combines vendor dependencies licenses for a project that manages dependencies with [dep](https://golang.github.io/dep).  
It fetches packages list from `Gopkg.lock` and searches their licenses in the corresponding `vendor/` subfolders.  
The retrieved metadata is processed by a [mustache](https://mustache.github.io/mustache.5.html) template, and printed to standard output.  
lisc is inspired by [license-webpack-plugin](https://www.npmjs.com/package/license-webpack-plugin) project.

## Installation
```go get github.com/kogu-io/lisc```

## Quick start
* Set your current working directory to a Golang project folder using `dep`,
* Run ```lisc > licenses.txt```,
* Use `licenses.txt` as you see fit.

## Specification
lisc requires no mandatory arguments to be passed.  
However, it allows you to control its behavior with the following options:
```
lisc [--source|-s project/path/]         # project root path
     [--names|-n license-file-name.txt]  # extra expected license file names
     [--template|-t template.mustache]   # template file path
     [--ensure|-e]                       # fail on any missing license
     [--help|-h]                         # print command help
```

## Template customization
lisc allows you to control the output format by passing a path to custom Mustache template file.  
In case not specified, lisc uses default template that is available on GitHub as [template.mustache](https://github.com/kogu-io/lisc/blob/master/template.mustache).  
License metadata objects contain the following tags:
* `Package` — package identifier (i.e. "github.com/kogu/lisc"),
* `Version` — package version (might be unspecified),
* `License` — package license text (if not found, defaults to "MIT"),
* `Explicit` — boolean flag indicating whether the license is explicitly specified and extracted.
