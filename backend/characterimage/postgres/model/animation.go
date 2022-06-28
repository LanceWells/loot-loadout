package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


Table: animation
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] body_type_id                                   INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] display_name                                   VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
[ 3] part_type                                      USER_DEFINED         null: true   primary: false  isArray: false  auto: false  col: USER_DEFINED    len: -1      default: []


JSON Sample
-------------------------------------
{    "id": 1,    "body_type_id": 67,    "display_name": "oehZbEIGUiKInMgMsSvoUJwvH",    "part_type": "DeShdtrAtJWocYnNZcmpCNoKC"}



*/

// Animation struct is a row record of the animation table in the charimage database
type Animation struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;column:id;type:INT4;" json:"id"`
	//[ 1] body_type_id                                   INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	BodyTypeID sql.NullInt64 `gorm:"column:body_type_id;type:INT4;" json:"body_type_id"`
	//[ 2] display_name                                   VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	DisplayName string `gorm:"column:display_name;type:VARCHAR;" json:"display_name"`
	//[ 3] part_type                                      USER_DEFINED         null: true   primary: false  isArray: false  auto: false  col: USER_DEFINED    len: -1      default: []
	PartType sql.NullString `gorm:"column:part_type;type:VARCHAR;" json:"part_type"`
}

var animationTableInfo = &TableInfo{
	Name: "animation",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "body_type_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "BodyTypeID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "body_type_id",
			ProtobufFieldName:  "body_type_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "display_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       -1,
			GoFieldName:        "DisplayName",
			GoFieldType:        "string",
			JSONFieldName:      "display_name",
			ProtobufFieldName:  "display_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "part_type",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "USER_DEFINED",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "USER_DEFINED",
			ColumnLength:       -1,
			GoFieldName:        "PartType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "part_type",
			ProtobufFieldName:  "part_type",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *Animation) TableName() string {
	return "animation"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Animation) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Animation) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Animation) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *Animation) TableInfo() *TableInfo {
	return animationTableInfo
}
