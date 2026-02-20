declare namespace kintone.types {
  interface CustomerFields {
    顧客情報メモ欄: kintone.fieldTypes.MultiLineText;
    郵便番号: kintone.fieldTypes.SingleLineText;
    Webサイト: kintone.fieldTypes.Link;
    建物名: kintone.fieldTypes.SingleLineText;
    業種: kintone.fieldTypes.DropDown;
    顧客ランク: kintone.fieldTypes.RadioButton;
    都道府県: kintone.fieldTypes.DropDown;
    支払日: kintone.fieldTypes.DropDown;
    住所: kintone.fieldTypes.SingleLineText;
    締め日: kintone.fieldTypes.DropDown;
    電話番号: kintone.fieldTypes.Link;
    FAX: kintone.fieldTypes.Link;
    会社名: kintone.fieldTypes.SingleLineText;
  }
  interface SavedCustomerFields extends CustomerFields {
    $id: kintone.fieldTypes.Id;
    $revision: kintone.fieldTypes.Revision;
    更新者: kintone.fieldTypes.Modifier;
    作成者: kintone.fieldTypes.Creator;
    顧客No: kintone.fieldTypes.RecordNumber;
    更新日時: kintone.fieldTypes.UpdatedTime;
    作成日時: kintone.fieldTypes.CreatedTime;
  }
}
