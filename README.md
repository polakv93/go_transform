# go_transform

A tiny command‑line utility written in Go that applies a JSON "transform" to one or more target JSON files. The transform file declares a special key `"__targets"` listing the files to modify; the rest of the JSON is merged into each target. Results overwrite the original target files.

Note: This README documents the current behavior as implemented in the repository. TODO items are included where information is unknown or where improvements could be made.

## Overview

- Language/Stack: Go (Go modules)
- Module name: `go_transform` (see `go.mod`)
- Entry point: `main.go`
- Primary package: `transform` (in `transform/transform.go`)
- Package manager/build tool: Go toolchain (no external dependencies)

How it works:
1. You provide a path to a JSON transform file.
2. The transform file contains:
   - A special `"__targets"` array of file paths to JSON files to update.
   - Any additional keys/values that will be merged into each target JSON file.
3. Each target JSON is read, merged (deep merge for nested objects), and written back to the same file.

Example transform file:

```json
{
  "__targets": [
    "exampleFiles/target1.json",
    "exampleFiles/target2.json"
  ],
  "somekey": "somevalue2"
}
```

## Requirements

- Go 1.25 or newer (per `go.mod`). The project was authored against `go 1.25`.
- A shell/terminal on your OS (Windows, macOS, or Linux). This repository is confirmed on Windows paths; forward slashes in JSON paths are also accepted by Go on Windows.

## Setup

No special setup is required beyond having Go installed.

- Verify Go installation:
  - `go version`
  - `go env GOPATH`

- Pull dependencies (none external):
  - `go mod tidy` (optional; `go.mod` currently has no third‑party requirements)

## Build and Run

You can run directly without building a binary:

- Run:
  - `go run . <path-to-transform.json>`
  - Example: `go run . exampleFiles/transform.json`

Or build a binary first:

- Build:
  - On Windows: `go build -o go_transform.exe .`
  - On macOS/Linux: `go build -o go_transform .`

- Run the built binary:
  - Windows: `./go_transform.exe exampleFiles/transform.json`
  - macOS/Linux: `./go_transform exampleFiles/transform.json`

## Install from a repository URL
`go install https://github.com/polakv93/go_transform@latest`  
After installation, the binary (go_transform) will be placed in GOPATH/bin or GOBIN if set.

### CLI help

- `go run . -help`
- Or with a built binary: `./go_transform -help`

Usage output:
```
Usage: [flags] <transform file path>
Available flags:
  -help
        Display help
```