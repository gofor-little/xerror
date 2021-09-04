## A package for formatted error stack traces

![GitHub tag (latest SemVer pre-release)](https://img.shields.io/github/v/tag/gofor-little/xerror?include_prereleases)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gofor-little/xerror)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://raw.githubusercontent.com/gofor-little/xerror/main/LICENSE)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/gofor-little/xerror/CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofor-little/xerror)](https://goreportcard.com/report/github.com/gofor-little/xerror)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gofor-little/xerror)](https://pkg.go.dev/github.com/gofor-little/xerror)

### Introduction
* Formatted error stack traces
* Supports JSON marshaling
* No dependencies outside the standard library

### Example
```go
package main

import (
	"fmt"
	"os"

	"github.com/gofor-little/xerror"
)

func main() {
	if err := RunApplication(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("application successfully started")
}

func RunApplication() error {
	if err := Initialize(); err != nil {
		return xerror.Wrap("failed to run application", err)
	}

	return nil
}

func Initialize() error {
	if err := LoadConfig(); err != nil {
		return xerror.Wrap("failed to initialize application", err)
	}

	return nil
}

func LoadConfig() error {
	_, err := os.Open("config.json")
	return xerror.Wrap("failed to load config", err)
}
```

Running the above will output...
```
main.RunApplication
        /home/ubuntu/xerror/main.go:21: failed to run application
main.Initialize
        /home/ubuntu/xerror/main.go:29: failed to initialize application
main.LoadConfig
        /home/ubuntu/xerror/main.go:37: failed to load config: open config.json: no such file or directory
exit status 1
```

Or can be marshaled into JSON and output...
```json
{
    "error": {
        "error": {
            "error": "open config.json: no such file or directory",
            "functionName": "main.LoadConfig",
            "fileName": "/home/ubuntu/xerror/cmd/main.go",
            "lineNumber": "39",
            "message": "failed to load config"
        },
        "functionName": "main.Initialize",
        "fileName": "/home/ubuntu/xerror/cmd/main.go",
        "lineNumber": "31",
        "message": "failed to initialize application"
    },
    "functionName": "main.RunApplication",
    "fileName": "/home/ubuntu/xerror/cmd/main.go",
    "lineNumber": "23",
    "message": "failed to run application"
}
```
### Testing
Run ```go test -v ./...``` in the root directory.
