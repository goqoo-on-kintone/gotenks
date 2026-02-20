# gotenks

kintone の TypeScript 型定義ファイル (.d.ts) から Go の型定義を生成する CLI ツール。

## Installation

```bash
go install github.com/goqoo-on-kintone/gotenks/cmd/gotenks@latest
```

## Usage

```bash
# ディレクトリ内の全 .d.ts ファイルを変換
gotenks -input ./dts -output ./gen/kintone

# 単一ファイルを変換
gotenks -input ./customer-fields.d.ts -output ./customer.go
```

## 入力ファイルの生成

入力となる `.d.ts` ファイルは [@kintone/dts-gen](https://github.com/kintone/js-sdk/tree/main/packages/dts-gen) で生成できます。

```bash
npx @kintone/dts-gen --base-url https://YOUR_SUBDOMAIN.cybozu.com \
  --app-id YOUR_APP_ID \
  --oauth-token YOUR_TOKEN \
  -o fields.d.ts
```

## Development

```bash
# ビルド
make build

# テスト
make test

# 開発用ビルド & 実行
make run
```

## License

MIT License - see [LICENSE](LICENSE) for details.
