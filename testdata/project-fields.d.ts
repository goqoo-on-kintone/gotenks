declare namespace kintone.types {
  interface ProjectFields {
    提案商品: kintone.fieldTypes.DropDown;
    アクション内容: kintone.fieldTypes.SingleLineText;
    詳細: kintone.fieldTypes.MultiLineText;
    案件名: kintone.fieldTypes.SingleLineText;
    売上: kintone.fieldTypes.Number;
    会社名: kintone.fieldTypes.SingleLineText;
    商談フェーズ: kintone.fieldTypes.DropDown;
    次回アクション日: kintone.fieldTypes.Date;
    確度: kintone.fieldTypes.DropDown;
    受注予定日: kintone.fieldTypes.Date;
    顧客No_: kintone.fieldTypes.Number;
    初回商談日: kintone.fieldTypes.Date;

    アクション担当者: kintone.fieldTypes.UserSelect;
    主担当: kintone.fieldTypes.UserSelect;
    主担当組織: kintone.fieldTypes.OrganizationSelect;
    契約書_申込書: kintone.fieldTypes.File;
  }
  interface SavedProjectFields extends ProjectFields {
    $id: kintone.fieldTypes.Id;
    $revision: kintone.fieldTypes.Revision;
    更新者: kintone.fieldTypes.Modifier;
    作成者: kintone.fieldTypes.Creator;
    案件No_: kintone.fieldTypes.RecordNumber;
    更新日時: kintone.fieldTypes.UpdatedTime;
    作成日時: kintone.fieldTypes.CreatedTime;
  }
}
