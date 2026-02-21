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

	sb.WriteString(fmt.Sprintf("// %s は kintone アプリのレコード型\n", structName))
	sb.WriteString(fmt.Sprintf("type %s struct {\n", structName))

	// extends がある場合、親の埋め込みではなく親のフィールドも展開する
	if iface.Extends != "" {
		// 親インターフェースを探す
		for _, parent := range result.Interfaces {
			if parent.Name == iface.Extends {
				for _, field := range parent.Fields {
					generateField(sb, field)
				}
				break
			}
		}
	}

	// 自身のフィールドを出力
	for _, field := range iface.Fields {
		generateField(sb, field)
	}

	sb.WriteString("}\n\n")
}

// generateField は1つのフィールドを出力する
func generateField(sb *strings.Builder, field parser.Field) {
	goFieldName := toGoIdentifier(field.Name)
	goTypeName := parser.TypeScriptToGoType(field.TypeName)

	// JSON タグで元のフィールド名を保持
	sb.WriteString(fmt.Sprintf("\t%s types.%s `json:\"%s\"`\n",
		goFieldName, goTypeName, field.Name))
}

// toGoIdentifier は日本語フィールド名を Go の有効な識別子に変換する
func toGoIdentifier(name string) string {
	// $id, $revision は特別扱い
	if name == "$id" {
		return "ID"
	}
	if name == "$revision" {
		return "Revision"
	}

	var result strings.Builder
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
