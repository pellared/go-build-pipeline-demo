# Go Build Pipeline Demo

This repository demonstrates a very basic build pipeline for a Go project created in:
- **[Make](./make)**
- **[Mage](./mage)**
- **[taskflow](./taskflow)**

## Build Pipeline

The build pipeline, called `all`, consists of following steps:

### `clean`

Removes files created during build. It logs which files are removed and which files it failed to remove.

### `fmt`

Simply runs `go fmt ./...`.

### `test`

Runs tests and generates coverage even if some test fails.
