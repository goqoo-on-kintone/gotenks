package kintone

import "github.com/goqoo-on-kintone/gotenks/internal/types"

type DetailsRow struct {
	ProductName types.SingleLineTextField `json:"Product_name"`
	Quantity types.NumberField `json:"Quantity"`
	UnitPrice types.NumberField `json:"Unit_price"`
}

type AllFieldsEn struct {
	SingleLineText types.SingleLineTextField `json:"Single_line_text"`
	MultiLineText types.MultiLineTextField `json:"Multi_line_text"`
	RichText types.RichTextField `json:"Rich_text"`
	Number types.NumberField `json:"Number"`
	Link types.LinkField `json:"Link"`
	Date types.DateField `json:"Date"`
	Time types.TimeField `json:"Time"`
	Datetime types.DateTimeField `json:"Datetime"`
	Dropdown types.DropDownField `json:"Dropdown"`
	RadioButton types.RadioButtonField `json:"Radio_button"`
	Checkbox types.CheckBoxField `json:"Checkbox"`
	MultiSelect types.MultiSelectField `json:"Multi_select"`
	Calc types.CalcField `json:"Calc"`
	UserSelect types.UserSelectField `json:"User_select"`
	OrganizationSelect types.OrganizationSelectField `json:"Organization_select"`
	GroupSelect types.GroupSelectField `json:"Group_select"`
	Attachment types.FileField `json:"Attachment"`
	Details types.Subtable[DetailsRow] `json:"Details"`
}

type SavedAllFieldsEn struct {
	AllFieldsEn
	ID types.IDField `json:"$id"`
	Revision types.RevisionField `json:"$revision"`
	CreatedBy types.CreatorField `json:"Created_by"`
	UpdatedBy types.ModifierField `json:"Updated_by"`
	UpdatedDatetime types.UpdatedTimeField `json:"Updated_datetime"`
	CreatedDatetime types.CreatedTimeField `json:"Created_datetime"`
	RecordNumber types.RecordNumberField `json:"Record_number"`
}

