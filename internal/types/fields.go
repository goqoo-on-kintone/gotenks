// Package types は kintone フィールドの Go 型定義を提供する
package types

// Entity はユーザー・組織・グループの共通構造
type Entity struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// FileInfo は添付ファイルの情報
type FileInfo struct {
	ContentType string `json:"contentType"`
	FileKey     string `json:"fileKey"`
	Name        string `json:"name"`
	Size        string `json:"size"`
}

// =============================================================================
// 文字列系フィールド
// =============================================================================

// SingleLineTextField は一行テキストフィールド
type SingleLineTextField struct {
	Type     string `json:"type,omitempty"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled,omitempty"`
	Error    string `json:"error,omitempty"`
}

// MultiLineTextField は複数行テキストフィールド
type MultiLineTextField struct {
	Type     string `json:"type,omitempty"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled,omitempty"`
	Error    string `json:"error,omitempty"`
}

// RichTextField はリッチエディタフィールド
type RichTextField struct {
	Type     string `json:"type,omitempty"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled,omitempty"`
	Error    string `json:"error,omitempty"`
}

// NumberField は数値フィールド
type NumberField struct {
	Type     string `json:"type,omitempty"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled,omitempty"`
	Error    string `json:"error,omitempty"`
}

// LinkField はリンクフィールド
type LinkField struct {
	Type     string `json:"type,omitempty"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled,omitempty"`
	Error    string `json:"error,omitempty"`
}

// =============================================================================
// 日時系フィールド
// =============================================================================

// DateField は日付フィールド
type DateField struct {
	Type     string `json:"type,omitempty"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled,omitempty"`
	Error    string `json:"error,omitempty"`
}

// TimeField は時刻フィールド
type TimeField struct {
	Type     string `json:"type,omitempty"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled,omitempty"`
	Error    string `json:"error,omitempty"`
}

// DateTimeField は日時フィールド
type DateTimeField struct {
	Type     string `json:"type,omitempty"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled,omitempty"`
	Error    string `json:"error,omitempty"`
}

// =============================================================================
// 選択系フィールド（単一選択）
// =============================================================================

// DropDownField はドロップダウンフィールド
type DropDownField struct {
	Type     string `json:"type,omitempty"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled,omitempty"`
	Error    string `json:"error,omitempty"`
}

// RadioButtonField はラジオボタンフィールド
type RadioButtonField struct {
	Type     string `json:"type,omitempty"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled,omitempty"`
	Error    string `json:"error,omitempty"`
}

// =============================================================================
// 選択系フィールド（複数選択）
// =============================================================================

// CheckBoxField はチェックボックスフィールド
type CheckBoxField struct {
	Type     string   `json:"type,omitempty"`
	Value    []string `json:"value"`
	Disabled bool     `json:"disabled,omitempty"`
	Error    string   `json:"error,omitempty"`
}

// MultiSelectField は複数選択フィールド
type MultiSelectField struct {
	Type     string   `json:"type,omitempty"`
	Value    []string `json:"value"`
	Disabled bool     `json:"disabled,omitempty"`
	Error    string   `json:"error,omitempty"`
}

// =============================================================================
// 計算フィールド
// =============================================================================

// CalcField は計算フィールド
type CalcField struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Disabled bool   `json:"disabled,omitempty"`
}

// =============================================================================
// ユーザー・組織・グループ選択フィールド
// =============================================================================

// UserSelectField はユーザー選択フィールド
type UserSelectField struct {
	Type     string   `json:"type,omitempty"`
	Value    []Entity `json:"value"`
	Disabled bool     `json:"disabled,omitempty"`
	Error    string   `json:"error,omitempty"`
}

// OrganizationSelectField は組織選択フィールド
type OrganizationSelectField struct {
	Type     string   `json:"type,omitempty"`
	Value    []Entity `json:"value"`
	Disabled bool     `json:"disabled,omitempty"`
	Error    string   `json:"error,omitempty"`
}

// GroupSelectField はグループ選択フィールド
type GroupSelectField struct {
	Type     string   `json:"type,omitempty"`
	Value    []Entity `json:"value"`
	Disabled bool     `json:"disabled,omitempty"`
	Error    string   `json:"error,omitempty"`
}

// =============================================================================
// ユーザー系フィールド（単一・システム）
// =============================================================================

// CreatorField は作成者フィールド
type CreatorField struct {
	Type  string `json:"type"`
	Value Entity `json:"value"`
}

// ModifierField は更新者フィールド
type ModifierField struct {
	Type  string `json:"type"`
	Value Entity `json:"value"`
}

// =============================================================================
// ファイルフィールド
// =============================================================================

// FileField は添付ファイルフィールド
type FileField struct {
	Type     string     `json:"type"`
	Value    []FileInfo `json:"value"`
	Disabled bool       `json:"disabled,omitempty"`
	Error    string     `json:"error,omitempty"`
}

// =============================================================================
// システムフィールド
// =============================================================================

// IDField はレコードIDフィールド
type IDField struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// RevisionField はリビジョンフィールド
type RevisionField struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// RecordNumberField はレコード番号フィールド
type RecordNumberField struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// CreatedTimeField は作成日時フィールド
type CreatedTimeField struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// UpdatedTimeField は更新日時フィールド
type UpdatedTimeField struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// =============================================================================
// サブテーブル
// =============================================================================

// SubtableRow はサブテーブルの1行を表す
type SubtableRow[T any] struct {
	ID    string `json:"id"`
	Value T      `json:"value"`
}

// Subtable はサブテーブルフィールド
type Subtable[T any] struct {
	Type  string           `json:"type"`
	Value []SubtableRow[T] `json:"value"`
}
