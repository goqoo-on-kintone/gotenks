package kintone

import "github.com/goqoo-on-kintone/gotenks/types"

type KDetailsRow struct {
	KProductName types.SingleLineTextField `json:"Product_name"`
	KQuantity    types.NumberField         `json:"Quantity"`
	KUnitPrice   types.NumberField         `json:"Unit_price"`
}

type AllFieldsEn struct {
	KSingleLineText     types.SingleLineTextField     `json:"Single_line_text"`
	KMultiLineText      types.MultiLineTextField      `json:"Multi_line_text"`
	KRichText           types.RichTextField           `json:"Rich_text"`
	KNumber             types.NumberField             `json:"Number"`
	KLink               types.LinkField               `json:"Link"`
	KDate               types.DateField               `json:"Date"`
	KTime               types.TimeField               `json:"Time"`
	KDatetime           types.DateTimeField           `json:"Datetime"`
	KDropdown           types.DropDownField           `json:"Dropdown"`
	KRadioButton        types.RadioButtonField        `json:"Radio_button"`
	KCheckbox           types.CheckBoxField           `json:"Checkbox"`
	KMultiSelect        types.MultiSelectField        `json:"Multi_select"`
	KCalc               types.CalcField               `json:"Calc"`
	KUserSelect         types.UserSelectField         `json:"User_select"`
	KOrganizationSelect types.OrganizationSelectField `json:"Organization_select"`
	KGroupSelect        types.GroupSelectField        `json:"Group_select"`
	KAttachment         types.FileField               `json:"Attachment"`
	KDetails            types.Subtable[KDetailsRow]   `json:"Details"`
}

type SavedAllFieldsEn struct {
	AllFieldsEn
	KID              types.IDField           `json:"$id"`
	KRevision        types.RevisionField     `json:"$revision"`
	KCreatedBy       types.CreatorField      `json:"Created_by"`
	KUpdatedBy       types.ModifierField     `json:"Updated_by"`
	KUpdatedDatetime types.UpdatedTimeField  `json:"Updated_datetime"`
	KCreatedDatetime types.CreatedTimeField  `json:"Created_datetime"`
	KRecordNumber    types.RecordNumberField `json:"Record_number"`
}
