# gotenks

CLI tool to generate Go type definitions from kintone TypeScript .d.ts files.

## Installation

```bash
npm install @goqoo/gotenks
```

Or use npx without installation:

```bash
npx @goqoo/gotenks -input ./fields.d.ts -output ./gen/
```

## Usage

```bash
# Convert all .d.ts files in a directory
gotenks -input ./dts -output ./gen/kintone

# Convert a single file
gotenks -input ./customer-fields.d.ts -output ./customer.go

# Specify package name
gotenks -input ./dts -output ./gen -package myapp
```

## Options

| Option | Default | Description |
|--------|---------|-------------|
| `-input` | (required) | Input .d.ts file or directory |
| `-output` | stdout | Output .go file or directory |
| `-package` | `kintone` | Package name for generated Go code |
| `-prefix` | `K` | Prefix for field names |

## Using Generated Code

Generated Go code imports the `types` package. Add it as a dependency:

```bash
go get github.com/goqoo-on-kintone/gotenks/types@latest
```

## More Information

See the [GitHub repository](https://github.com/goqoo-on-kintone/gotenks) for full documentation.

## License

MIT
