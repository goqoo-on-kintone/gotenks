# gotenks

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

[English](/README.md) | 日本語

kintone の TypeScript 型定義ファイル (.d.ts) から Go の型定義を生成する CLI ツール。

## インストール

### Homebrew（macOS / Linux）

```bash
brew install goqoo-on-kintone/tap/gotenks
```

### go install

```bash
go install github.com/goqoo-on-kintone/gotenks/cmd/gotenks@latest
```

### バイナリダウンロード

[GitHub Releases](https://github.com/goqoo-on-kintone/gotenks/releases) からプラットフォーム別のバイナリをダウンロード：

- `gotenks_X.X.X_darwin_amd64.tar.gz`（macOS Intel）
- `gotenks_X.X.X_darwin_arm64.tar.gz`（macOS Apple Silicon）
- `gotenks_X.X.X_linux_amd64.tar.gz`（Linux x64）
- `gotenks_X.X.X_linux_arm64.tar.gz`（Linux ARM64）
- `gotenks_X.X.X_windows_amd64.zip`（Windows x64）
- `gotenks_X.X.X_windows_arm64.zip`（Windows ARM64）

### ソースからビルド

```bash
git clone https://github.com/goqoo-on-kintone/gotenks.git
cd gotenks
make build
# bin/gotenks が生成される
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

## 対応フィールド型

| カテゴリ | TypeScript | Go 型 |
|---------|------------|-------|
| **文字列系** | `SingleLineText` | `SingleLineTextField` |
| | `MultiLineText` | `MultiLineTextField` |
| | `RichText` | `RichTextField` |
| | `Number` | `NumberField` |
| | `Link` | `LinkField` |
| **日時系** | `Date` | `DateField` |
| | `Time` | `TimeField` |
| | `DateTime` | `DateTimeField` |
| **選択系（単一）** | `DropDown` | `DropDownField` |
| | `RadioButton` | `RadioButtonField` |
| **選択系（複数）** | `CheckBox` | `CheckBoxField` |
| | `MultiSelect` | `MultiSelectField` |
| **計算** | `Calc` | `CalcField` |
| **ユーザー/組織/グループ** | `UserSelect` | `UserSelectField` |
| | `OrganizationSelect` | `OrganizationSelectField` |
| | `GroupSelect` | `GroupSelectField` |
| | `Creator` | `CreatorField` |
| | `Modifier` | `ModifierField` |
| **ファイル** | `File` | `FileField` |
| **システム** | `Id` | `IDField` |
| | `Revision` | `RevisionField` |
| | `RecordNumber` | `RecordNumberField` |
| | `CreatedTime` | `CreatedTimeField` |
| | `UpdatedTime` | `UpdatedTimeField` |
| **サブテーブル** | `SUBTABLE` | `Subtable[T]` |

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
