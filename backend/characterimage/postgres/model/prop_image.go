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


Table: prop_image
[ 0] prop_id                                        INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] x                                              INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
[ 2] y                                              INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
[ 3] image_bytes                                    BYTEA                null: true   primary: false  isArray: false  auto: false  col: BYTEA           len: -1      default: []


JSON Sample
-------------------------------------
{    "prop_id": 30,    "x": 15,    "y": 75,    "image_bytes": "elvDaDaPxiPFZQpcSKcPeUVsY"}



*/

// PropImage struct is a row record of the prop_image table in the charimage database
type PropImage struct {
	//[ 0] prop_id                                        INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	PropID int32 `gorm:"primary_key;column:prop_id;type:INT4;" json:"prop_id"`
	//[ 1] x                                              INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
	X sql.NullInt64 `gorm:"column:x;type:INT2;" json:"x"`
	//[ 2] y                                              INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
	Y sql.NullInt64 `gorm:"column:y;type:INT2;" json:"y"`
	//[ 3] image_bytes                                    BYTEA                null: true   primary: false  isArray: false  auto: false  col: BYTEA           len: -1      default: []
	ImageBytes sql.NullString `gorm:"column:image_bytes;type:BYTEA;" json:"image_bytes"`
}

var prop_imageTableInfo = &TableInfo{
	Name: "prop_image",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "prop_id",
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
			GoFieldName:        "PropID",
			GoFieldType:        "int32",
			JSONFieldName:      "prop_id",
			ProtobufFieldName:  "prop_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "x",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT2",
			DatabaseTypePretty: "INT2",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT2",
			ColumnLength:       -1,
			GoFieldName:        "X",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "x",
			ProtobufFieldName:  "x",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "y",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT2",
			DatabaseTypePretty: "INT2",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT2",
			ColumnLength:       -1,
			GoFieldName:        "Y",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "y",
			ProtobufFieldName:  "y",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "image_bytes",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "BYTEA",
			DatabaseTypePretty: "BYTEA",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BYTEA",
			ColumnLength:       -1,
			GoFieldName:        "ImageBytes",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "image_bytes",
			ProtobufFieldName:  "image_bytes",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *PropImage) TableName() string {
	return "prop_image"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PropImage) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PropImage) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PropImage) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PropImage) TableInfo() *TableInfo {
	return prop_imageTableInfo
}
