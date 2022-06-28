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


Table: dynamic_part_pixel
[ 0] color_string_id                                INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] dynamic_part_id                                INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] x                                              INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
[ 3] y                                              INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []


JSON Sample
-------------------------------------
{    "color_string_id": 6,    "dynamic_part_id": 42,    "x": 91,    "y": 71}



*/

// DynamicPartPixel struct is a row record of the dynamic_part_pixel table in the charimage database
type DynamicPartPixel struct {
	//[ 0] color_string_id                                INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	ColorStringID int32 `gorm:"primary_key;column:color_string_id;type:INT4;" json:"color_string_id"`
	//[ 1] dynamic_part_id                                INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	DynamicPartID int32 `gorm:"primary_key;column:dynamic_part_id;type:INT4;" json:"dynamic_part_id"`
	//[ 2] x                                              INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
	X sql.NullInt64 `gorm:"column:x;type:INT2;" json:"x"`
	//[ 3] y                                              INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
	Y sql.NullInt64 `gorm:"column:y;type:INT2;" json:"y"`
}

var dynamic_part_pixelTableInfo = &TableInfo{
	Name: "dynamic_part_pixel",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "color_string_id",
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
			GoFieldName:        "ColorStringID",
			GoFieldType:        "int32",
			JSONFieldName:      "color_string_id",
			ProtobufFieldName:  "color_string_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "dynamic_part_id",
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
			GoFieldName:        "DynamicPartID",
			GoFieldType:        "int32",
			JSONFieldName:      "dynamic_part_id",
			ProtobufFieldName:  "dynamic_part_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *DynamicPartPixel) TableName() string {
	return "dynamic_part_pixel"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *DynamicPartPixel) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *DynamicPartPixel) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *DynamicPartPixel) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *DynamicPartPixel) TableInfo() *TableInfo {
	return dynamic_part_pixelTableInfo
}
