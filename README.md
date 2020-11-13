# Go Build Pipeline Demo

[Presentation](https://docs.google.com/presentation/d/1fJ26B1D1VkxC-1DppegPCe8YOaA3Ayrbke2yAx-Kzcs/edit?usp=sharing).

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
