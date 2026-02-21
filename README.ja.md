# gotenks

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

[English](/README.md) | 日本語

kintone の TypeScript 型定義ファイル (.d.ts) から Go の型定義を生成する CLI ツール。

## インストール

```bash
go install github.com/goqoo-on-kintone/gotenks/cmd/gotenks@latest
```

## 使い方

```bash
# ディレクトリ内の全 .d.ts ファイルを変換
gotenks -input ./dts -output ./gen/kintone

# 単一ファイルを変換
gotenks -input ./customer-fields.d.ts -output ./customer.go

# パッケージ名を指定
gotenks -input ./dts -output ./gen -package myapp

# prefix なしで生成（ASCII フィールドコードのみの場合）
gotenks -input ./dts -output ./gen -prefix ""
```

## オプション

| オプション | デフォルト | 説明 |
|-----------|-----------|------|
| `-input` | (必須) | 入力 .d.ts ファイルまたはディレクトリ |
| `-output` | stdout | 出力 .go ファイルまたはディレクトリ |
| `-package` | `kintone` | 生成する Go コードのパッケージ名 |
| `-prefix` | `K` | フィールド名のプレフィックス（日本語フィールド名のエクスポート用） |

## 入力ファイルの生成

入力となる `.d.ts` ファイルは [@kintone/dts-gen](https://github.com/kintone/js-sdk/tree/main/packages/dts-gen) で生成できます。

```bash
npx @kintone/dts-gen --base-url https://YOUR_SUBDOMAIN.cybozu.com \
  --app-id YOUR_APP_ID \
  --oauth-token YOUR_TOKEN \
  -o fields.d.ts
```

## 開発

```bash
# ビルド
make build

# テスト
make test

# ビルド & 実行
make run
```

## ライセンス

MIT License - 詳細は [LICENSE](LICENSE) を参照してください。
