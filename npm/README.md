# gotenks

CLI tool to generate Go type definitions from kintone TypeScript .d.ts files.

## Installation

```bash
npm install @goqoo/gotenks
```

Or use npx without installation:

```bash
npx @goqoo/gotenks ./fields.d.ts -o ./gen/
```

## Usage

```bash
# Convert all .d.ts files in a directory
gotenks ./dts -o ./gen/kintone

# Convert a single file
gotenks ./customer-fields.d.ts -o ./customer.go

# Specify package name
gotenks ./dts -o ./gen -package myapp
```

## Options

| Option | Short | Default | Description |
|--------|-------|---------|-------------|
| `-output` | `-o` | stdout | Output .go file or directory |
| `-package` | | `kintone` | Package name for generated Go code |
| `-prefix` | | `K` | Prefix for field names |

## Using Generated Code

Generated Go code imports the `types` package. Add it as a dependency:

```bash
go get github.com/goqoo-on-kintone/gotenks/types@latest
```

## More Information

See the [GitHub repository](https://github.com/goqoo-on-kintone/gotenks) for full documentation.

## License

MIT
