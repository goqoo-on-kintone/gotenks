# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## プロジェクト概要

gotenks は、kintone アプリの TypeScript 型定義ファイル (.d.ts) から Go の型定義を生成する CLI ツール。

## ディレクトリ構成

```
├── cmd/gotenks/         # CLI エントリーポイント
├── internal/
│   ├── parser/          # .d.ts パーサー
│   ├── generator/       # Go コード生成
│   └── types/           # kintone フィールド型定義
├── testdata/            # テスト用 .d.ts ファイル
└── examples/            # 入出力サンプル
```

## 開発コマンド

```bash
make build      # ビルド → bin/gotenks
make test       # テスト実行
make run        # ビルド & 実行
make run-test   # testdata/ を使って実行
make lint       # go vet + staticcheck
make fmt        # go fmt
```

## 参考リポジトリ

- [@kintone/dts-gen](https://github.com/kintone/js-sdk/tree/main/packages/dts-gen) - kintone 型定義生成ツールの公式実装。パーサーや型定義の実装を参照する際に有用

## kintone 型定義の構造（入力ファイル）

```typescript
// testdata/*.d.ts の基本構造
declare namespace kintone.types {
  interface XxxFields {
    フィールド名: kintone.fieldTypes.SingleLineText;
    // ...
  }
  interface SavedXxxFields extends XxxFields {
    $id: kintone.fieldTypes.Id;
    // システムフィールド
  }
}
```

## コーディング規約

- コミュニケーション・コードコメントは日本語
- コンソール出力は状況に応じて英語/日本語を使い分け
