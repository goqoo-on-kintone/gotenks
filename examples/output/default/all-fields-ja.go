package kintone

import "github.com/goqoo-on-kintone/gotenks/types"

type K明細Row struct {
	K品名 types.SingleLineTextField `json:"品名"`
	K数量 types.NumberField         `json:"数量"`
	K単価 types.NumberField         `json:"単価"`
}

type AllFieldsJa struct {
	K一行テキスト   types.SingleLineTextField     `json:"一行テキスト"`
	K複数行テキスト  types.MultiLineTextField      `json:"複数行テキスト"`
	Kリッチエディタ  types.RichTextField           `json:"リッチエディタ"`
	K数値       types.NumberField             `json:"数値"`
	Kリンク      types.LinkField               `json:"リンク"`
	K日付       types.DateField               `json:"日付"`
	K時刻       types.TimeField               `json:"時刻"`
	K日時       types.DateTimeField           `json:"日時"`
	Kドロップダウン  types.DropDownField           `json:"ドロップダウン"`
	Kラジオボタン   types.RadioButtonField        `json:"ラジオボタン"`
	Kチェックボックス types.CheckBoxField           `json:"チェックボックス"`
	K複数選択     types.MultiSelectField        `json:"複数選択"`
	K計算       types.CalcField               `json:"計算"`
	Kユーザー選択   types.UserSelectField         `json:"ユーザー選択"`
	K組織選択     types.OrganizationSelectField `json:"組織選択"`
	Kグループ選択   types.GroupSelectField        `json:"グループ選択"`
	K添付ファイル   types.FileField               `json:"添付ファイル"`
	K明細       types.Subtable[K明細Row]        `json:"明細"`
}

type SavedAllFieldsJa struct {
	AllFieldsJa
	KID       types.IDField           `json:"$id"`
	KRevision types.RevisionField     `json:"$revision"`
	K更新者      types.ModifierField     `json:"更新者"`
	K作成者      types.CreatorField      `json:"作成者"`
	Kレコード番号   types.RecordNumberField `json:"レコード番号"`
	K更新日時     types.UpdatedTimeField  `json:"更新日時"`
	K作成日時     types.CreatedTimeField  `json:"作成日時"`
}
