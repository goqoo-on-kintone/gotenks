// Package parser は .d.ts ファイルをパースして型情報を抽出する
package parser

import (
	"os"
	"regexp"
	"strings"
)

// Field は kintone フィールドの情報を表す
type Field struct {
	Name           string  // フィールド名（日本語など）
	TypeName       string  // kintone.fieldTypes.Xxx の Xxx 部分（通常フィールドの場合）
	IsSubtable     bool    // サブテーブルかどうか
	SubtableFields []Field // サブテーブル内のフィールド（IsSubtable=true の場合のみ）
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

	// interface を検索して、それぞれのボディを抽出
	interfaceRe := regexp.MustCompile(`interface\s+(\w+)(?:\s+extends\s+(\w+))?\s*\{`)

	// 全ての interface 開始位置を見つける
	matches := interfaceRe.FindAllStringSubmatchIndex(content, -1)
	for _, match := range matches {
		name := content[match[2]:match[3]]
		var extends string
		if match[4] != -1 {
			extends = content[match[4]:match[5]]
		}

		// interface の開始 { の位置
		braceStart := match[1] - 1

		// 対応する } を見つける
		body := extractBraceContent(content, braceStart)

		iface := Interface{
			Name:    name,
			Extends: extends,
			Fields:  parseFields(body),
		}
		result.Interfaces = append(result.Interfaces, iface)
	}

	return result, nil
}

// extractBraceContent は開始位置の { から対応する } までの内容を抽出する
func extractBraceContent(content string, start int) string {
	if start >= len(content) || content[start] != '{' {
		return ""
	}

	depth := 0
	for i := start; i < len(content); i++ {
		switch content[i] {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				// { と } を除いた中身を返す
				return content[start+1 : i]
			}
		}
	}
	return ""
}

// parseFields はインターフェース内のフィールド定義をパースする
func parseFields(content string) []Field {
	var fields []Field

	// フィールドを1つずつ解析する
	// 行ごとに処理しつつ、SUBTABLE のようなネスト構造も考慮
	i := 0
	for i < len(content) {
		// 空白をスキップ
		for i < len(content) && (content[i] == ' ' || content[i] == '\t' || content[i] == '\n' || content[i] == '\r') {
			i++
		}
		if i >= len(content) {
			break
		}

		// コメント行をスキップ（// で始まる行）
		if i+1 < len(content) && content[i] == '/' && content[i+1] == '/' {
			// 行末までスキップ
			for i < len(content) && content[i] != '\n' {
				i++
			}
			continue
		}

		// フィールド名を抽出（: まで）
		fieldStart := i
		for i < len(content) && content[i] != ':' {
			i++
		}
		if i >= len(content) {
			break
		}
		fieldName := strings.TrimSpace(content[fieldStart:i])
		i++ // ':' をスキップ

		// 空白をスキップ
		for i < len(content) && (content[i] == ' ' || content[i] == '\t') {
			i++
		}

		// 型の部分を解析
		if i < len(content) && content[i] == '{' {
			// オブジェクトリテラル（SUBTABLE など）
			braceContent := extractBraceContent(content, i)
			field := parseSubtableField(fieldName, braceContent)
			if field != nil {
				fields = append(fields, *field)
			}
			// 対応する } の次へ移動
			i += len(braceContent) + 2 // +2 for { and }
			// ; をスキップ
			for i < len(content) && content[i] != ';' {
				i++
			}
			i++
		} else if strings.HasPrefix(content[i:], "kintone.fieldTypes.") {
			// 通常のフィールド
			typeStart := i + len("kintone.fieldTypes.")
			typeEnd := typeStart
			for typeEnd < len(content) && content[typeEnd] != ';' {
				typeEnd++
			}
			typeName := strings.TrimSpace(content[typeStart:typeEnd])
			fields = append(fields, Field{
				Name:     fieldName,
				TypeName: typeName,
			})
			i = typeEnd + 1
		} else {
			// その他（スキップ）
			for i < len(content) && content[i] != ';' && content[i] != '\n' {
				i++
			}
			i++
		}
	}

	return fields
}

// parseSubtableField は SUBTABLE フィールドをパースする
func parseSubtableField(name, content string) *Field {
	// type: "SUBTABLE" を確認
	if !strings.Contains(content, `type: "SUBTABLE"`) && !strings.Contains(content, `type: 'SUBTABLE'`) {
		return nil
	}

	// value の中の value オブジェクトを探す
	// value: { id: string; value: { ... } }[]
	// 内側の value: { ... } からフィールドを抽出

	// 内側の value: { を探す
	innerValueRe := regexp.MustCompile(`value:\s*\{[^}]*value:\s*\{`)
	match := innerValueRe.FindStringIndex(content)
	if match == nil {
		return nil
	}

	// 内側の value の { の位置を見つける
	// "value: {" を2回探す
	firstValue := strings.Index(content, "value:")
	if firstValue == -1 {
		return nil
	}
	// 最初の value の { を見つける
	firstBrace := strings.Index(content[firstValue:], "{")
	if firstBrace == -1 {
		return nil
	}
	firstBrace += firstValue

	// その中で2番目の value: { を探す
	innerContent := extractBraceContent(content, firstBrace)
	secondValue := strings.Index(innerContent, "value:")
	if secondValue == -1 {
		return nil
	}
	secondBrace := strings.Index(innerContent[secondValue:], "{")
	if secondBrace == -1 {
		return nil
	}
	secondBrace += secondValue

	// 内側のフィールド定義を抽出
	innerFieldsContent := extractBraceContent(innerContent, secondBrace)

	// 内側のフィールドをパース（再帰的に parseFields を使用）
	innerFields := parseFields(innerFieldsContent)

	return &Field{
		Name:           name,
		IsSubtable:     true,
		SubtableFields: innerFields,
	}
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
