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


Table: static_part_image
[ 0] static_part_id                                 INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] x                                              INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
[ 2] y                                              INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
[ 3] image_bytes                                    BYTEA                null: true   primary: false  isArray: false  auto: false  col: BYTEA           len: -1      default: []


JSON Sample
-------------------------------------
{    "static_part_id": 11,    "x": 21,    "y": 11,    "image_bytes": "PHkeXAQwyTVxOuVkULbGkPWxV"}



*/

// StaticPartImage struct is a row record of the static_part_image table in the charimage database
type StaticPartImage struct {
	//[ 0] static_part_id                                 INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	StaticPartID int32 `gorm:"primary_key;column:static_part_id;type:INT4;" json:"static_part_id"`
	//[ 1] x                                              INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
	X sql.NullInt64 `gorm:"column:x;type:INT2;" json:"x"`
	//[ 2] y                                              INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
	Y sql.NullInt64 `gorm:"column:y;type:INT2;" json:"y"`
	//[ 3] image_bytes                                    BYTEA                null: true   primary: false  isArray: false  auto: false  col: BYTEA           len: -1      default: []
	ImageBytes sql.NullString `gorm:"column:image_bytes;type:BYTEA;" json:"image_bytes"`
}

var static_part_imageTableInfo = &TableInfo{
	Name: "static_part_image",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "static_part_id",
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
			GoFieldName:        "StaticPartID",
			GoFieldType:        "int32",
			JSONFieldName:      "static_part_id",
			ProtobufFieldName:  "static_part_id",
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
func (s *StaticPartImage) TableName() string {
	return "static_part_image"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *StaticPartImage) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *StaticPartImage) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *StaticPartImage) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *StaticPartImage) TableInfo() *TableInfo {
	return static_part_imageTableInfo
}
