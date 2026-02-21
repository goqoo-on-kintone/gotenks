package parser

import (
	"testing"
)

func TestParse_SimpleInterface(t *testing.T) {
	input := `
declare namespace kintone.types {
  interface CustomerFields {
    会社名: kintone.fieldTypes.SingleLineText;
    電話番号: kintone.fieldTypes.Link;
  }
}
`
	result, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(result.Interfaces) != 1 {
		t.Fatalf("expected 1 interface, got %d", len(result.Interfaces))
	}

	iface := result.Interfaces[0]
	if iface.Name != "CustomerFields" {
		t.Errorf("expected name CustomerFields, got %s", iface.Name)
	}
	if iface.Extends != "" {
		t.Errorf("expected no extends, got %s", iface.Extends)
	}
	if len(iface.Fields) != 2 {
		t.Errorf("expected 2 fields, got %d", len(iface.Fields))
	}
}

func TestParse_InterfaceWithExtends(t *testing.T) {
	input := `
declare namespace kintone.types {
  interface CustomerFields {
    会社名: kintone.fieldTypes.SingleLineText;
  }
  interface SavedCustomerFields extends CustomerFields {
    $id: kintone.fieldTypes.Id;
    $revision: kintone.fieldTypes.Revision;
  }
}
`
	result, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(result.Interfaces) != 2 {
		t.Fatalf("expected 2 interfaces, got %d", len(result.Interfaces))
	}

	// SavedCustomerFields
	saved := result.Interfaces[1]
	if saved.Name != "SavedCustomerFields" {
		t.Errorf("expected name SavedCustomerFields, got %s", saved.Name)
	}
	if saved.Extends != "CustomerFields" {
		t.Errorf("expected extends CustomerFields, got %s", saved.Extends)
	}
	if len(saved.Fields) != 2 {
		t.Errorf("expected 2 fields, got %d", len(saved.Fields))
	}
}

func TestParse_AllFieldTypes(t *testing.T) {
	input := `
declare namespace kintone.types {
  interface TestFields {
    text: kintone.fieldTypes.SingleLineText;
    multiText: kintone.fieldTypes.MultiLineText;
    richText: kintone.fieldTypes.RichText;
    number: kintone.fieldTypes.Number;
    calc: kintone.fieldTypes.Calc;
    date: kintone.fieldTypes.Date;
    time: kintone.fieldTypes.Time;
    dateTime: kintone.fieldTypes.DateTime;
    link: kintone.fieldTypes.Link;
    dropdown: kintone.fieldTypes.DropDown;
    radio: kintone.fieldTypes.RadioButton;
    checkbox: kintone.fieldTypes.CheckBox;
    multiSelect: kintone.fieldTypes.MultiSelect;
    user: kintone.fieldTypes.UserSelect;
    org: kintone.fieldTypes.OrganizationSelect;
    group: kintone.fieldTypes.GroupSelect;
    file: kintone.fieldTypes.File;
    creator: kintone.fieldTypes.Creator;
    modifier: kintone.fieldTypes.Modifier;
    id: kintone.fieldTypes.Id;
    revision: kintone.fieldTypes.Revision;
    recordNumber: kintone.fieldTypes.RecordNumber;
    createdTime: kintone.fieldTypes.CreatedTime;
    updatedTime: kintone.fieldTypes.UpdatedTime;
  }
}
`
	result, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(result.Interfaces) != 1 {
		t.Fatalf("expected 1 interface, got %d", len(result.Interfaces))
	}

	iface := result.Interfaces[0]
	if len(iface.Fields) != 24 {
		t.Errorf("expected 24 fields, got %d", len(iface.Fields))
	}
}

func TestParse_Subtable(t *testing.T) {
	input := `
declare namespace kintone.types {
  interface FacilityFields {
    施設名: kintone.fieldTypes.SingleLineText;
    部屋タイプ: {
      type: "SUBTABLE";
      value: {
        id: string;
        value: {
          部屋タイプ名: kintone.fieldTypes.SingleLineText;
          定員: kintone.fieldTypes.Number;
        };
      }[];
    };
    住所: kintone.fieldTypes.SingleLineText;
  }
}
`
	result, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(result.Interfaces) != 1 {
		t.Fatalf("expected 1 interface, got %d", len(result.Interfaces))
	}

	iface := result.Interfaces[0]
	// 施設名, 部屋タイプ (SUBTABLE), 住所 = 3 fields
	if len(iface.Fields) != 3 {
		t.Errorf("expected 3 fields, got %d", len(iface.Fields))
		for i, f := range iface.Fields {
			t.Logf("field %d: %+v", i, f)
		}
	}

	// 部屋タイプがSUBTABLEとして認識されているか
	var subtableField *Field
	for _, f := range iface.Fields {
		if f.Name == "部屋タイプ" {
			subtableField = &f
			break
		}
	}

	if subtableField == nil {
		t.Fatal("部屋タイプ field not found")
	}

	if !subtableField.IsSubtable {
		t.Error("部屋タイプ should be a subtable")
	}

	if len(subtableField.SubtableFields) != 2 {
		t.Errorf("expected 2 subtable fields, got %d", len(subtableField.SubtableFields))
	}
}

func TestParseFields(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Field
	}{
		{
			name:  "single field",
			input: "会社名: kintone.fieldTypes.SingleLineText;",
			expected: []Field{
				{Name: "会社名", TypeName: "SingleLineText"},
			},
		},
		{
			name:  "multiple fields",
			input: "name: kintone.fieldTypes.SingleLineText;\n    email: kintone.fieldTypes.Link;",
			expected: []Field{
				{Name: "name", TypeName: "SingleLineText"},
				{Name: "email", TypeName: "Link"},
			},
		},
		{
			name:  "field with special characters",
			input: "$id: kintone.fieldTypes.Id;",
			expected: []Field{
				{Name: "$id", TypeName: "Id"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := parseFields(tt.input)
			if len(fields) != len(tt.expected) {
				t.Fatalf("expected %d fields, got %d", len(tt.expected), len(fields))
			}
			for i, f := range fields {
				if f.Name != tt.expected[i].Name {
					t.Errorf("field %d: expected name %s, got %s", i, tt.expected[i].Name, f.Name)
				}
				if f.TypeName != tt.expected[i].TypeName {
					t.Errorf("field %d: expected type %s, got %s", i, tt.expected[i].TypeName, f.TypeName)
				}
			}
		})
	}
}

func TestTypeScriptToGoType(t *testing.T) {
	tests := []struct {
		tsType   string
		expected string
	}{
		{"SingleLineText", "SingleLineTextField"},
		{"MultiLineText", "MultiLineTextField"},
		{"RichText", "RichTextField"},
		{"Number", "NumberField"},
		{"Calc", "CalcField"},
		{"Date", "DateField"},
		{"Time", "TimeField"},
		{"DateTime", "DateTimeField"},
		{"Link", "LinkField"},
		{"DropDown", "DropDownField"},
		{"RadioButton", "RadioButtonField"},
		{"CheckBox", "CheckBoxField"},
		{"MultiSelect", "MultiSelectField"},
		{"UserSelect", "UserSelectField"},
		{"OrganizationSelect", "OrganizationSelectField"},
		{"GroupSelect", "GroupSelectField"},
		{"File", "FileField"},
		{"Creator", "CreatorField"},
		{"Modifier", "ModifierField"},
		{"Id", "IDField"},
		{"Revision", "RevisionField"},
		{"RecordNumber", "RecordNumberField"},
		{"CreatedTime", "CreatedTimeField"},
		{"UpdatedTime", "UpdatedTimeField"},
		// 未知の型
		{"Unknown", "UnknownField"},
	}

	for _, tt := range tests {
		t.Run(tt.tsType, func(t *testing.T) {
			got := TypeScriptToGoType(tt.tsType)
			if got != tt.expected {
				t.Errorf("TypeScriptToGoType(%s) = %s, want %s", tt.tsType, got, tt.expected)
			}
		})
	}
}

func TestParseFile(t *testing.T) {
	result, err := ParseFile("../../testdata/customer-fields.d.ts")
	if err != nil {
		t.Fatalf("ParseFile failed: %v", err)
	}

	if len(result.Interfaces) != 2 {
		t.Fatalf("expected 2 interfaces, got %d", len(result.Interfaces))
	}

	// CustomerFields
	if result.Interfaces[0].Name != "CustomerFields" {
		t.Errorf("expected first interface to be CustomerFields, got %s", result.Interfaces[0].Name)
	}

	// SavedCustomerFields
	if result.Interfaces[1].Name != "SavedCustomerFields" {
		t.Errorf("expected second interface to be SavedCustomerFields, got %s", result.Interfaces[1].Name)
	}
	if result.Interfaces[1].Extends != "CustomerFields" {
		t.Errorf("expected SavedCustomerFields to extend CustomerFields, got %s", result.Interfaces[1].Extends)
	}
}
