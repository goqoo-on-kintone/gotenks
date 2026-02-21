declare namespace kintone.types {
  interface AllFieldsJa {
    // 文字列系
    一行テキスト: kintone.fieldTypes.SingleLineText;
    複数行テキスト: kintone.fieldTypes.MultiLineText;
    リッチエディタ: kintone.fieldTypes.RichText;
    数値: kintone.fieldTypes.Number;
    リンク: kintone.fieldTypes.Link;

    // 日時系
    日付: kintone.fieldTypes.Date;
    時刻: kintone.fieldTypes.Time;
    日時: kintone.fieldTypes.DateTime;

    // 選択系（単一）
    ドロップダウン: kintone.fieldTypes.DropDown;
    ラジオボタン: kintone.fieldTypes.RadioButton;

    // 選択系（複数）
    チェックボックス: kintone.fieldTypes.CheckBox;
    複数選択: kintone.fieldTypes.MultiSelect;

    // 計算
    計算: kintone.fieldTypes.Calc;

    // ユーザー/組織/グループ選択
    ユーザー選択: kintone.fieldTypes.UserSelect;
    組織選択: kintone.fieldTypes.OrganizationSelect;
    グループ選択: kintone.fieldTypes.GroupSelect;

    // ファイル
    添付ファイル: kintone.fieldTypes.File;

    // サブテーブル
    明細: {
      type: "SUBTABLE";
      value: {
        id: string;
        value: {
          品名: kintone.fieldTypes.SingleLineText;
          数量: kintone.fieldTypes.Number;
          単価: kintone.fieldTypes.Number;
        };
      }[];
    };
  }
  interface SavedAllFieldsJa extends AllFieldsJa {
    $id: kintone.fieldTypes.Id;
    $revision: kintone.fieldTypes.Revision;
    更新者: kintone.fieldTypes.Modifier;
    作成者: kintone.fieldTypes.Creator;
    レコード番号: kintone.fieldTypes.RecordNumber;
    更新日時: kintone.fieldTypes.UpdatedTime;
    作成日時: kintone.fieldTypes.CreatedTime;
  }
}
