// Package generator はパースした型情報から Go コードを生成する
package generator

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/goqoo-on-kintone/gotenks/internal/parser"
)

// Config は生成オプション
type Config struct {
	PackageName string // 生成するパッケージ名
}

// Generate はパース結果から Go コードを生成する
func Generate(result *parser.ParseResult, config Config) string {
	var sb strings.Builder

	// パッケージ宣言
	sb.WriteString(fmt.Sprintf("package %s\n\n", config.PackageName))

	// types パッケージのインポート
	sb.WriteString("import \"github.com/goqoo-on-kintone/gotenks/internal/types\"\n\n")

	// 各インターフェースを struct に変換
	for _, iface := range result.Interfaces {
		generateStruct(&sb, iface, result)
	}

	return sb.String()
}

// generateStruct は interface を Go の struct として出力する
func generateStruct(sb *strings.Builder, iface parser.Interface, result *parser.ParseResult) {
	structName := iface.Name

	// まずサブテーブルの Row 構造体を先に生成
	// 注: 親の Row 構造体は親インターフェース処理時に既に生成されているため、
	// extends の場合は自身のフィールドのみ処理する
	generateSubtableRowStructs(sb, iface.Fields)

	sb.WriteString(fmt.Sprintf("type %s struct {\n", structName))

	// extends がある場合、親の構造体を埋め込む
	if iface.Extends != "" {
		sb.WriteString(fmt.Sprintf("\t%s\n", iface.Extends))
	}

	// 自身のフィールドを出力
	for _, field := range iface.Fields {
		generateField(sb, field)
	}

	sb.WriteString("}\n\n")
}

// generateSubtableRowStructs はサブテーブルの Row 構造体を生成する
func generateSubtableRowStructs(sb *strings.Builder, fields []parser.Field) {
	for _, field := range fields {
		if !field.IsSubtable {
			continue
		}

		rowStructName := toGoIdentifier(field.Name) + "Row"

		sb.WriteString(fmt.Sprintf("type %s struct {\n", rowStructName))

		for _, subField := range field.SubtableFields {
			generateField(sb, subField)
		}

		sb.WriteString("}\n\n")
	}
}

// generateField は1つのフィールドを出力する
func generateField(sb *strings.Builder, field parser.Field) {
	goFieldName := toGoIdentifier(field.Name)

	if field.IsSubtable {
		// サブテーブルフィールド
		rowStructName := goFieldName + "Row"
		sb.WriteString(fmt.Sprintf("\t%s types.Subtable[%s] `json:\"%s\"`\n",
			goFieldName, rowStructName, field.Name))
	} else {
		// 通常フィールド
		goTypeName := parser.TypeScriptToGoType(field.TypeName)
		sb.WriteString(fmt.Sprintf("\t%s types.%s `json:\"%s\"`\n",
			goFieldName, goTypeName, field.Name))
	}
}

// toGoIdentifier はフィールド名を Go の有効な識別子に変換する
// すべてのフィールドに K プレフィックスを付けてエクスポート可能にする
func toGoIdentifier(name string) string {
	var result strings.Builder
	result.WriteString("K") // Kintone の K プレフィックス

	// $id, $revision は特別扱い
	if name == "$id" {
		result.WriteString("ID")
		return result.String()
	}
	if name == "$revision" {
		result.WriteString("Revision")
		return result.String()
	}

	capitalizeNext := true

	for _, r := range name {
		switch {
		case r == '_' || r == '-' || r == '/' || r == ' ':
			// 区切り文字は次の文字を大文字にする
			capitalizeNext = true
		case unicode.IsLetter(r) || unicode.IsNumber(r):
			if capitalizeNext {
				result.WriteRune(unicode.ToUpper(r))
				capitalizeNext = false
			} else {
				result.WriteRune(r)
			}
		default:
			// その他の文字（日本語など）はそのまま
			result.WriteRune(r)
			capitalizeNext = false
		}
	}

	return result.String()
}
