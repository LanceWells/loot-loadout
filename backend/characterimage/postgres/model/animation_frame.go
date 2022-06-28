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


Table: animation_frame
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] animation_id                                   INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] frame_index                                    INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": 89,    "animation_id": 60,    "frame_index": 96}



*/

// AnimationFrame struct is a row record of the animation_frame table in the charimage database
type AnimationFrame struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;column:id;type:INT4;" json:"id"`
	//[ 1] animation_id                                   INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	AnimationID sql.NullInt64 `gorm:"column:animation_id;type:INT4;" json:"animation_id"`
	//[ 2] frame_index                                    INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	FrameIndex int32 `gorm:"column:frame_index;type:INT4;" json:"frame_index"`
}

var animation_frameTableInfo = &TableInfo{
	Name: "animation_frame",
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
			Name:               "animation_id",
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
			GoFieldName:        "AnimationID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "animation_id",
			ProtobufFieldName:  "animation_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "frame_index",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "FrameIndex",
			GoFieldType:        "int32",
			JSONFieldName:      "frame_index",
			ProtobufFieldName:  "frame_index",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AnimationFrame) TableName() string {
	return "animation_frame"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AnimationFrame) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AnimationFrame) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AnimationFrame) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AnimationFrame) TableInfo() *TableInfo {
	return animation_frameTableInfo
}
