package generator

import (
	"strings"
	"testing"

	"github.com/goqoo-on-kintone/gotenks/internal/parser"
)

func TestGenerate_SimpleStruct(t *testing.T) {
	result := &parser.ParseResult{
		Interfaces: []parser.Interface{
			{
				Name: "CustomerFields",
				Fields: []parser.Field{
					{Name: "会社名", TypeName: "SingleLineText"},
					{Name: "電話番号", TypeName: "Link"},
				},
			},
		},
	}

	config := Config{PackageName: "kintone"}
	code := Generate(result, config)

	// パッケージ宣言のチェック
	if !strings.Contains(code, "package kintone") {
		t.Error("generated code should contain package declaration")
	}

	// import のチェック
	if !strings.Contains(code, `import "github.com/goqoo-on-kintone/gotenks/internal/types"`) {
		t.Error("generated code should import types package")
	}

	// struct のチェック
	if !strings.Contains(code, "type CustomerFields struct") {
		t.Error("generated code should contain CustomerFields struct")
	}

	// フィールドのチェック（Kプレフィックス付き）
	if !strings.Contains(code, "K会社名 types.SingleLineTextField") {
		t.Error("generated code should contain K会社名 field")
	}
	if !strings.Contains(code, `json:"会社名"`) {
		t.Error("generated code should contain json tag for 会社名")
	}
}

func TestGenerate_WithExtends(t *testing.T) {
	result := &parser.ParseResult{
		Interfaces: []parser.Interface{
			{
				Name: "CustomerFields",
				Fields: []parser.Field{
					{Name: "会社名", TypeName: "SingleLineText"},
				},
			},
			{
				Name:    "SavedCustomerFields",
				Extends: "CustomerFields",
				Fields: []parser.Field{
					{Name: "$id", TypeName: "Id"},
					{Name: "$revision", TypeName: "Revision"},
				},
			},
		},
	}

	config := Config{PackageName: "kintone"}
	code := Generate(result, config)

	// SavedCustomerFields struct の定義があるか
	if !strings.Contains(code, "type SavedCustomerFields struct") {
		t.Error("generated code should contain SavedCustomerFields struct")
	}

	// SavedCustomerFields の定義部分を抽出
	savedStart := strings.Index(code, "type SavedCustomerFields struct")
	if savedStart == -1 {
		t.Fatal("SavedCustomerFields not found")
	}
	savedEnd := strings.Index(code[savedStart:], "\n}\n")
	savedStruct := code[savedStart : savedStart+savedEnd]

	// 親の構造体が埋め込まれているか
	if !strings.Contains(savedStruct, "CustomerFields") {
		t.Error("SavedCustomerFields should embed CustomerFields")
	}

	// 自身のフィールドが含まれているか（Kプレフィックス付き）
	if !strings.Contains(savedStruct, "KID types.IDField") {
		t.Error("SavedCustomerFields should contain KID field")
	}
}

func TestToGoIdentifier(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"$id", "KID"},
		{"$revision", "KRevision"},
		{"会社名", "K会社名"},
		{"user_name", "KUserName"},
		{"user-name", "KUserName"},
		{"user/name", "KUserName"},
		{"user name", "KUserName"},
		{"firstName", "KFirstName"},
		{"FAX", "KFAX"},
		{"Webサイト", "KWebサイト"},
		{"施設No__タップ施設ID", "K施設Noタップ施設ID"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := toGoIdentifier(tt.input)
			if got != tt.expected {
				t.Errorf("toGoIdentifier(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestGenerate_AllFieldTypes(t *testing.T) {
	result := &parser.ParseResult{
		Interfaces: []parser.Interface{
			{
				Name: "TestFields",
				Fields: []parser.Field{
					{Name: "text", TypeName: "SingleLineText"},
					{Name: "number", TypeName: "Number"},
					{Name: "date", TypeName: "Date"},
					{Name: "user", TypeName: "UserSelect"},
					{Name: "file", TypeName: "File"},
					{Name: "creator", TypeName: "Creator"},
				},
			},
		},
	}

	config := Config{PackageName: "kintone"}
	code := Generate(result, config)

	// 各フィールド型が正しく変換されているか
	expectations := []string{
		"types.SingleLineTextField",
		"types.NumberField",
		"types.DateField",
		"types.UserSelectField",
		"types.FileField",
		"types.CreatorField",
	}

	for _, exp := range expectations {
		if !strings.Contains(code, exp) {
			t.Errorf("generated code should contain %s", exp)
		}
	}
}

func TestGenerate_EmptyResult(t *testing.T) {
	result := &parser.ParseResult{
		Interfaces: []parser.Interface{},
	}

	config := Config{PackageName: "kintone"}
	code := Generate(result, config)

	// パッケージ宣言は含まれるべき
	if !strings.Contains(code, "package kintone") {
		t.Error("generated code should contain package declaration even with empty result")
	}
}

func TestGenerate_CustomPackageName(t *testing.T) {
	result := &parser.ParseResult{
		Interfaces: []parser.Interface{
			{
				Name:   "TestFields",
				Fields: []parser.Field{},
			},
		},
	}

	config := Config{PackageName: "mypackage"}
	code := Generate(result, config)

	if !strings.Contains(code, "package mypackage") {
		t.Error("generated code should use custom package name")
	}
}
