# lisc - LIcense SCanner

lisc is naiive implementation for scanning and outputting licenses from used vendor-packages in golang projects that **use dep (https://golang.github.io/dep/) for dependency management**. It scans packages located in vendor subfolder that are referenced in Gopkg.lock file and writes the content from LICENSE files to stdout.

lisc is inspired by license-webpack-plugin (https://www.npmjs.com/package/license-webpack-plugin).

### Installation
```go get github.com/kogu-io/lisc```

### Usage
```lisc --help```

```lisc -t path/to/my/template.mustache.file -s path/to/my/source/folder```

### Templating
lisc uses Mustache (https://mustache.github.io/mustache.5.html) templates. Example template is available in GitHub as [template.mustache](https://github.com/kogu-io/lisc/blob/master/template.mustache). 

Following tags are available:
* Package - name of the package (Example: github.com/kogu/lisc)
* Version - version of the package (if provided)
* License - the text of license from the content of the package. If this is omitted lisc outputs default "MIT" license

### Example
* Change to project folder
* Create template for output (or copy sample from github) as ```template.mustache```
* Run ```lisc |tee licenses.txt```
* Use licenses.txt as you want.
