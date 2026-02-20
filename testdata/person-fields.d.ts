declare namespace kintone.types {
  interface PersonFields {
    備考: kintone.fieldTypes.MultiLineText;
    名: kintone.fieldTypes.SingleLineText;
    メールアドレス: kintone.fieldTypes.Link;
    役職: kintone.fieldTypes.SingleLineText;
    姓: kintone.fieldTypes.SingleLineText;
    顧客名: kintone.fieldTypes.SingleLineText;
    お名前: kintone.fieldTypes.SingleLineText;
    部署: kintone.fieldTypes.SingleLineText;
    電話番号: kintone.fieldTypes.Link;
    お名前_フリガナ: kintone.fieldTypes.SingleLineText;
    決裁権: kintone.fieldTypes.RadioButton;
    顧客No_: kintone.fieldTypes.Number;
    名_フリガナ: kintone.fieldTypes.SingleLineText;
    姓_フリガナ: kintone.fieldTypes.SingleLineText;
    携帯番号: kintone.fieldTypes.Link;

    添付ファイル_名刺等: kintone.fieldTypes.File;
  }
  interface SavedPersonFields extends PersonFields {
    $id: kintone.fieldTypes.Id;
    $revision: kintone.fieldTypes.Revision;
    更新者: kintone.fieldTypes.Modifier;
    作成者: kintone.fieldTypes.Creator;
    更新日時: kintone.fieldTypes.UpdatedTime;
    担当者No_: kintone.fieldTypes.RecordNumber;
    作成日時: kintone.fieldTypes.CreatedTime;
  }
}
