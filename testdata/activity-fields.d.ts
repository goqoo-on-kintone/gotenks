declare namespace kintone.types {
  interface ActivityFields {
    案件No: kintone.fieldTypes.Number;
    顧客No: kintone.fieldTypes.Number;
    タイトル: kintone.fieldTypes.SingleLineText;
    案件名: kintone.fieldTypes.SingleLineText;
    対応日付: kintone.fieldTypes.Date;
    対応種別: kintone.fieldTypes.DropDown;
    内容: kintone.fieldTypes.MultiLineText;
    会社名: kintone.fieldTypes.SingleLineText;

    対応者: kintone.fieldTypes.UserSelect;
    所属組織: kintone.fieldTypes.OrganizationSelect;
    添付ファイル: kintone.fieldTypes.File;
  }
  interface SavedActivityFields extends ActivityFields {
    $id: kintone.fieldTypes.Id;
    $revision: kintone.fieldTypes.Revision;
    更新者: kintone.fieldTypes.Modifier;
    作成者: kintone.fieldTypes.Creator;
    レコード番号: kintone.fieldTypes.RecordNumber;
    更新日時: kintone.fieldTypes.UpdatedTime;
    作成日時: kintone.fieldTypes.CreatedTime;
  }
}
