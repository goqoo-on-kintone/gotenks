# gotenks

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

English | [日本語](/README.ja.md)

A CLI tool that generates Go type definitions from kintone TypeScript type definition files (.d.ts).

## Installation

### Homebrew (macOS / Linux)

```bash
brew install goqoo-on-kintone/tap/gotenks
```

### go install

```bash
go install github.com/goqoo-on-kintone/gotenks/cmd/gotenks@latest
```

### Binary Download

Download platform-specific binaries from [GitHub Releases](https://github.com/goqoo-on-kintone/gotenks/releases):

- `gotenks_X.X.X_darwin_amd64.tar.gz` (macOS Intel)
- `gotenks_X.X.X_darwin_arm64.tar.gz` (macOS Apple Silicon)
- `gotenks_X.X.X_linux_amd64.tar.gz` (Linux x64)
- `gotenks_X.X.X_linux_arm64.tar.gz` (Linux ARM64)
- `gotenks_X.X.X_windows_amd64.zip` (Windows x64)
- `gotenks_X.X.X_windows_arm64.zip` (Windows ARM64)

### Build from Source

```bash
git clone https://github.com/goqoo-on-kintone/gotenks.git
cd gotenks
make build
# Binary is generated at bin/gotenks
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

## Using Generated Code

Generated Go code imports the `types` package. Add it as a dependency to your project:

```bash
go get github.com/goqoo-on-kintone/gotenks/types@latest
```

## Supported Field Types

| Category | TypeScript | Go Type |
|----------|------------|---------|
| **String** | `SingleLineText` | `SingleLineTextField` |
| | `MultiLineText` | `MultiLineTextField` |
| | `RichText` | `RichTextField` |
| | `Number` | `NumberField` |
| | `Link` | `LinkField` |
| **Date/Time** | `Date` | `DateField` |
| | `Time` | `TimeField` |
| | `DateTime` | `DateTimeField` |
| **Selection (single)** | `DropDown` | `DropDownField` |
| | `RadioButton` | `RadioButtonField` |
| **Selection (multiple)** | `CheckBox` | `CheckBoxField` |
| | `MultiSelect` | `MultiSelectField` |
| **Calculation** | `Calc` | `CalcField` |
| **User/Org/Group** | `UserSelect` | `UserSelectField` |
| | `OrganizationSelect` | `OrganizationSelectField` |
| | `GroupSelect` | `GroupSelectField` |
| | `Creator` | `CreatorField` |
| | `Modifier` | `ModifierField` |
| **File** | `File` | `FileField` |
| **System** | `Id` | `IDField` |
| | `Revision` | `RevisionField` |
| | `RecordNumber` | `RecordNumberField` |
| | `CreatedTime` | `CreatedTimeField` |
| | `UpdatedTime` | `UpdatedTimeField` |
| **Subtable** | `SUBTABLE` | `Subtable[T]` |

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
