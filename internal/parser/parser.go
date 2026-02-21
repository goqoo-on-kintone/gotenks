// Package parser は .d.ts ファイルをパースして型情報を抽出する
package parser

import (
	"os"
	"regexp"
	"strings"
)

// Field は kintone フィールドの情報を表す
type Field struct {
	Name     string // フィールド名（日本語など）
	TypeName string // kintone.fieldTypes.Xxx の Xxx 部分
}

// Interface は TypeScript の interface 定義を表す
type Interface struct {
	Name    string  // インターフェース名 (例: CustomerFields)
	Extends string  // 継承元 (例: CustomerFields、空なら継承なし)
	Fields  []Field // フィールド一覧
}

// ParseResult はパース結果を表す
type ParseResult struct {
	Interfaces []Interface
}

// ParseFile は .d.ts ファイルをパースする
func ParseFile(filepath string) (*ParseResult, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return Parse(string(content))
}

// Parse は .d.ts の内容をパースする
func Parse(content string) (*ParseResult, error) {
	result := &ParseResult{}

	// interface の正規表現
	// interface XxxFields { ... } または interface SavedXxxFields extends XxxFields { ... }
	interfaceRe := regexp.MustCompile(`interface\s+(\w+)(?:\s+extends\s+(\w+))?\s*\{([^}]*)\}`)

	matches := interfaceRe.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		iface := Interface{
			Name:    match[1],
			Extends: match[2],
			Fields:  parseFields(match[3]),
		}
		result.Interfaces = append(result.Interfaces, iface)
	}

	return result, nil
}

// parseFields はインターフェース内のフィールド定義をパースする
func parseFields(content string) []Field {
	var fields []Field

	// 通常のフィールド: フィールド名: kintone.fieldTypes.Xxx;
	fieldRe := regexp.MustCompile(`(\S+):\s*kintone\.fieldTypes\.(\w+);`)

	matches := fieldRe.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		fields = append(fields, Field{
			Name:     strings.TrimSpace(match[1]),
			TypeName: match[2],
		})
	}

	return fields
}

// TypeScriptToGoType は TypeScript のフィールド型を Go の型名に変換する
func TypeScriptToGoType(tsType string) string {
	mapping := map[string]string{
		// 文字列系
		"SingleLineText": "SingleLineTextField",
		"MultiLineText":  "MultiLineTextField",
		"RichText":       "RichTextField",
		"Number":         "NumberField",
		"Link":           "LinkField",
		// 日時系
		"Date":     "DateField",
		"Time":     "TimeField",
		"DateTime": "DateTimeField",
		// 選択系（単一）
		"DropDown":    "DropDownField",
		"RadioButton": "RadioButtonField",
		// 選択系（複数）
		"CheckBox":    "CheckBoxField",
		"MultiSelect": "MultiSelectField",
		// 計算
		"Calc": "CalcField",
		// ユーザー・組織・グループ
		"UserSelect":         "UserSelectField",
		"OrganizationSelect": "OrganizationSelectField",
		"GroupSelect":        "GroupSelectField",
		"Creator":            "CreatorField",
		"Modifier":           "ModifierField",
		// ファイル
		"File": "FileField",
		// システム
		"Id":           "IDField",
		"Revision":     "RevisionField",
		"RecordNumber": "RecordNumberField",
		"CreatedTime":  "CreatedTimeField",
		"UpdatedTime":  "UpdatedTimeField",
	}

	if goType, ok := mapping[tsType]; ok {
		return goType
	}
	return tsType + "Field" // 未知の型はそのまま + Field
}
