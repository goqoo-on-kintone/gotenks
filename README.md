# gotenks

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

English | [日本語](/README.ja.md)

A CLI tool that generates Go type definitions from kintone TypeScript type definition files (.d.ts).

## Installation

```bash
go install github.com/goqoo-on-kintone/gotenks/cmd/gotenks@latest
```

## Usage

```bash
# Convert all .d.ts files in a directory
gotenks -input ./dts -output ./gen/kintone

# Convert a single file
gotenks -input ./customer-fields.d.ts -output ./customer.go

# Specify package name
gotenks -input ./dts -output ./gen -package myapp

# Generate without prefix (for ASCII-only field codes)
gotenks -input ./dts -output ./gen -prefix ""
```

## Options

| Option | Default | Description |
|--------|---------|-------------|
| `-input` | (required) | Input .d.ts file or directory |
| `-output` | stdout | Output .go file or directory |
| `-package` | `kintone` | Package name for generated Go code |
| `-prefix` | `K` | Prefix for field names (for exporting Japanese field names) |

## Generating Input Files

Input `.d.ts` files can be generated using [@kintone/dts-gen](https://github.com/kintone/js-sdk/tree/main/packages/dts-gen).

```bash
npx @kintone/dts-gen --base-url https://YOUR_SUBDOMAIN.cybozu.com \
  --app-id YOUR_APP_ID \
  --oauth-token YOUR_TOKEN \
  -o fields.d.ts
```

## Development

```bash
# Build
make build

# Test
make test

# Build & Run
make run
```

## License

MIT License - see [LICENSE](LICENSE) for details.
