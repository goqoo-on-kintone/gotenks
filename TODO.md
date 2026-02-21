# dts → Go 型生成ツール TODO

dts/*.d.ts から Go の kintone 型定義を生成するツールの開発タスク。

---

## 完了済みタスク

- [x] プロジェクト構成を決める
- [x] kintone フィールド型の Go 定義を作成（全24種類対応）
- [x] TypeScript パーサーを実装（コメント行スキップ対応）
- [x] Go コード生成ロジックを実装（extends → 埋め込み対応）
- [x] サブテーブル型の変換に対応（Subtable[T] ジェネリクス）
- [x] フィールド名のエスケープ処理（日本語・記号対応）
- [x] CLI インターフェースを実装（-input, -output, -package, -prefix オプション）
- [x] テスト用 dts で動作確認

---

## 対応済みフィールド型

### 2. kintone フィールド型の Go 定義を作成

対応すべき型（@kintone/dts-gen の kintone.d.ts より）:

| TypeScript | Go 型 | 値の構造 |
|------------|-------|----------|
| **文字列系** |||
| `SingleLineText` | `SingleLineTextField` | `string` |
| `MultiLineText` | `MultiLineTextField` | `string` |
| `RichText` | `RichTextField` | `string` |
| `Number` | `NumberField` | `string` |
| `Link` | `LinkField` | `string` |
| **日時系** |||
| `Date` | `DateField` | `string` |
| `Time` | `TimeField` | `string` |
| `DateTime` | `DateTimeField` | `string` |
| **選択系（単一）** |||
| `DropDown` | `DropDownField` | `string` |
| `RadioButton` | `RadioButtonField` | `string` |
| **選択系（複数）** |||
| `CheckBox` | `CheckBoxField` | `[]string` |
| `MultiSelect` | `MultiSelectField` | `[]string` |
| **計算** |||
| `Calc` | `CalcField` | `string` |
| **ユーザー/組織系（配列）** |||
| `UserSelect` | `UserSelectField` | `[]Entity{Code,Name}` |
| `OrganizationSelect` | `OrganizationSelectField` | `[]Entity{Code,Name}` |
| `GroupSelect` | `GroupSelectField` | `[]Entity{Code,Name}` |
| **ユーザー系（単一）** |||
| `Creator` | `CreatorField` | `Entity{Code,Name}` |
| `Modifier` | `ModifierField` | `Entity{Code,Name}` |
| **ファイル** |||
| `File` | `FileField` | `[]FileInfo{ContentType,FileKey,Name,Size}` |
| **システム** |||
| `Id` | `IdField` | `string` |
| `Revision` | `RevisionField` | `string` |
| `RecordNumber` | `RecordNumberField` | `string` |
| `UpdatedTime` | `UpdatedTimeField` | `string` |
| `CreatedTime` | `CreatedTimeField` | `string` |
| **サブテーブル** |||
| `SUBTABLE` | `Subtable[T]` | ジェネリクス |

※ Lookup フィールドと関連レコード一覧は dts-gen では除外される（実体は参照先のフィールド型）

### 3. TypeScript パーサーを実装

- `.d.ts` ファイルを読み込み
- `interface XxxFields { ... }` を抽出
- フィールド名と型のマッピングを取得
- サブテーブルの入れ子構造を解析

### 4. Go コード生成ロジックを実装

- interface → struct への変換
- JSON タグの付与（フィールドコード保持）
- `SavedXxxFields` の extends 処理

### 5. サブテーブル型の変換に対応

```typescript
// 入力
"配属・異動": {
  type: "SUBTABLE";
  value: { id: string; value: { ... } }[];
}
```

```go
// 出力
配属異動 Subtable[配属異動Row] `json:"配属・異動"`

type 配属異動Row struct { ... }
```

### 6. フィールド名のエスケープ処理

- 日本語フィールド名 → Go 識別子への変換
- 記号（`__`, `-`, `/`）の処理
- JSON タグで元のフィールドコードを保持

### 7. CLI インターフェースを実装

```bash
gotenks -input ./dts -output ./gen/kintone
gotenks -input ./dts/facility-master-fields.d.ts  # 単一ファイル
```

### 8. テスト用 dts で動作確認

- 既存の dts ファイルで生成テスト
- 生成された Go コードがコンパイル通るか確認

---

## 判断ポイント

| 項目 | 選択肢 |
|------|--------|
| **パーサー実装** | 正規表現 vs 簡易 AST パーサー |
| **ツール言語** | Go 自体で書く vs TypeScript で書いて Go 出力 |
| **フィールド名** | 日本語のまま vs 英語変換 vs 両方（タグで元名保持） |

## @kintone/dts-gen のアーキテクチャ参考

```
src/
├── cli-parser.ts                    # CLI引数パース
├── converters/
│   └── fileldtype-converter.ts      # フィールド型をグループに分類
├── templates/
│   ├── converter.ts                 # FieldTypeGroups → TS式変換
│   ├── template.ts                  # ファイル出力（ESLint+Prettier整形）
│   └── expressions/                 # TypeScript AST的ビルダー
│       ├── namespace.ts
│       ├── typedefinitions.ts
│       └── fields.ts
└── kintone/clients/                 # kintone API クライアント
```

**設計のポイント:**
- フィールドを「グループ」に分類して処理（strings, calc, users, lists, subtables, file）
- ビルダーパターンで出力コードを組み立て
- Lookup / 関連レコード一覧は変換対象から除外

---

## 参考: dts ファイルの構造

```typescript
declare namespace kintone.types {
  interface FacilityMasterFields {
    施設略語コード: kintone.fieldTypes.SingleLineText;
    施設No__タップ施設ID: kintone.fieldTypes.Number;
    // ...
    部屋タイプ: {
      type: "SUBTABLE";
      value: {
        id: string;
        value: {
          部屋タイプ_テーブル: kintone.fieldTypes.SingleLineText;
        };
      }[];
    };
  }
  interface SavedFacilityMasterFields extends FacilityMasterFields {
    $id: kintone.fieldTypes.Id;
    $revision: kintone.fieldTypes.Revision;
    更新者: kintone.fieldTypes.Modifier;
    作成者: kintone.fieldTypes.Creator;
    レコード番号: kintone.fieldTypes.RecordNumber;
    更新日時: kintone.fieldTypes.UpdatedTime;
    作成日時: kintone.fieldTypes.CreatedTime;
  }
}
```
