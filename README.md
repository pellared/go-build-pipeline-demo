# Go Build Pipeline Demo

[![Demo](https://img.youtube.com/vi/AtiUf0uJ4YE/hqdefault.jpg)](https://www.youtube.com/watch?v=AtiUf0uJ4YE)

[Presentation](https://docs.google.com/presentation/d/1fJ26B1D1VkxC-1DppegPCe8YOaA3Ayrbke2yAx-Kzcs/edit?usp=sharing).

This repository demonstrates a very basic build pipeline for a Go project created in:
- **[Make](./make)**
- **[Mage](./mage)**
- **[taskflow](./taskflow)**

## Build Pipeline

The build pipeline, called `all`, consists of the following steps:

### `clean`

Removes files created during the build. It logs which files are removed and which files failed to be removed.

### `fmt`

Simply runs `go fmt ./...`. This target is imported as it can be reused by multiple projects.

### `test`

Runs tests and generates code coverage even if any test fails.
