declare namespace kintone.types {
  interface FacilityFields {
    施設名: kintone.fieldTypes.SingleLineText;
    施設コード: kintone.fieldTypes.SingleLineText;
    住所: kintone.fieldTypes.SingleLineText;
    電話番号: kintone.fieldTypes.Link;
    部屋タイプ: {
      type: "SUBTABLE";
      value: {
        id: string;
        value: {
          部屋タイプ名: kintone.fieldTypes.SingleLineText;
          定員: kintone.fieldTypes.Number;
          料金: kintone.fieldTypes.Number;
        };
      }[];
    };
    設備一覧: {
      type: "SUBTABLE";
      value: {
        id: string;
        value: {
          設備名: kintone.fieldTypes.SingleLineText;
          数量: kintone.fieldTypes.Number;
        };
      }[];
    };
  }
  interface SavedFacilityFields extends FacilityFields {
    $id: kintone.fieldTypes.Id;
    $revision: kintone.fieldTypes.Revision;
    更新者: kintone.fieldTypes.Modifier;
    作成者: kintone.fieldTypes.Creator;
    レコード番号: kintone.fieldTypes.RecordNumber;
    更新日時: kintone.fieldTypes.UpdatedTime;
    作成日時: kintone.fieldTypes.CreatedTime;
  }
}
