declare namespace kintone.types {
  interface AllFieldsEn {
    // String fields
    Single_line_text: kintone.fieldTypes.SingleLineText;
    Multi_line_text: kintone.fieldTypes.MultiLineText;
    Rich_text: kintone.fieldTypes.RichText;
    Number: kintone.fieldTypes.Number;
    Link: kintone.fieldTypes.Link;

    // Date/Time fields
    Date: kintone.fieldTypes.Date;
    Time: kintone.fieldTypes.Time;
    Datetime: kintone.fieldTypes.DateTime;

    // Selection fields (single)
    Dropdown: kintone.fieldTypes.DropDown;
    Radio_button: kintone.fieldTypes.RadioButton;

    // Selection fields (multiple)
    Checkbox: kintone.fieldTypes.CheckBox;
    Multi_select: kintone.fieldTypes.MultiSelect;

    // Calculation
    Calc: kintone.fieldTypes.Calc;

    // User/Organization/Group selection
    User_select: kintone.fieldTypes.UserSelect;
    Organization_select: kintone.fieldTypes.OrganizationSelect;
    Group_select: kintone.fieldTypes.GroupSelect;

    // File
    Attachment: kintone.fieldTypes.File;

    // Subtable
    Details: {
      type: "SUBTABLE";
      value: {
        id: string;
        value: {
          Product_name: kintone.fieldTypes.SingleLineText;
          Quantity: kintone.fieldTypes.Number;
          Unit_price: kintone.fieldTypes.Number;
        };
      }[];
    };
  }
  interface SavedAllFieldsEn extends AllFieldsEn {
    $id: kintone.fieldTypes.Id;
    $revision: kintone.fieldTypes.Revision;
    Created_by: kintone.fieldTypes.Creator;
    Updated_by: kintone.fieldTypes.Modifier;
    Updated_datetime: kintone.fieldTypes.UpdatedTime;
    Created_datetime: kintone.fieldTypes.CreatedTime;
    Record_number: kintone.fieldTypes.RecordNumber;
  }
}
