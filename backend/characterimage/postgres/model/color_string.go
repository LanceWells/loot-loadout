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


Table: color_string
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] hexstring                                      VARCHAR(9)           null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 9       default: []


JSON Sample
-------------------------------------
{    "id": 72,    "hexstring": "DVQiZUjuADwNTWdHooKQTLGSn"}



*/

// ColorString struct is a row record of the color_string table in the charimage database
type ColorString struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;column:id;type:INT4;" json:"id"`
	//[ 1] hexstring                                      VARCHAR(9)           null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 9       default: []
	Hexstring sql.NullString `gorm:"column:hexstring;type:VARCHAR;size:9;" json:"hexstring"`
}

var color_stringTableInfo = &TableInfo{
	Name: "color_string",
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
			Name:               "hexstring",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(9)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       9,
			GoFieldName:        "Hexstring",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "hexstring",
			ProtobufFieldName:  "hexstring",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *ColorString) TableName() string {
	return "color_string"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ColorString) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ColorString) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ColorString) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ColorString) TableInfo() *TableInfo {
	return color_stringTableInfo
}
