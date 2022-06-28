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


Table: body_type
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] display_name                                   VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []


JSON Sample
-------------------------------------
{    "id": 86,    "display_name": "dLFfrZUneBlSUKnLVirPuwTPy"}



*/

// BodyType struct is a row record of the body_type table in the charimage database
type BodyType struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;column:id;type:INT4;" json:"id"`
	//[ 1] display_name                                   VARCHAR              null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: -1      default: []
	DisplayName string `gorm:"column:display_name;type:VARCHAR;" json:"display_name"`
}

var body_typeTableInfo = &TableInfo{
	Name: "body_type",
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
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *BodyType) TableName() string {
	return "body_type"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BodyType) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BodyType) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BodyType) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BodyType) TableInfo() *TableInfo {
	return body_typeTableInfo
}
