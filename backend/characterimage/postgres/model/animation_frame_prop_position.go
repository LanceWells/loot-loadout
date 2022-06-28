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


Table: animation_frame_prop_position
[ 0] animation_frame_id                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] part_type                                      USER_DEFINED         null: false  primary: true   isArray: false  auto: false  col: USER_DEFINED    len: -1      default: []
[ 2] x                                              INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
[ 3] y                                              INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
[ 4] rotation                                       INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []


JSON Sample
-------------------------------------
{    "animation_frame_id": 58,    "part_type": "luufUfVDTqOEBepamdXPsAkRT",    "x": 64,    "y": 37,    "rotation": 97}



*/

// AnimationFramePropPosition struct is a row record of the animation_frame_prop_position table in the charimage database
type AnimationFramePropPosition struct {
	//[ 0] animation_frame_id                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	AnimationFrameID int32 `gorm:"primary_key;column:animation_frame_id;type:INT4;" json:"animation_frame_id"`
	//[ 1] part_type                                      USER_DEFINED         null: false  primary: true   isArray: false  auto: false  col: USER_DEFINED    len: -1      default: []
	PartType string `gorm:"primary_key;column:part_type;type:VARCHAR;" json:"part_type"`
	//[ 2] x                                              INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
	X sql.NullInt64 `gorm:"column:x;type:INT2;" json:"x"`
	//[ 3] y                                              INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
	Y sql.NullInt64 `gorm:"column:y;type:INT2;" json:"y"`
	//[ 4] rotation                                       INT2                 null: true   primary: false  isArray: false  auto: false  col: INT2            len: -1      default: []
	Rotation sql.NullInt64 `gorm:"column:rotation;type:INT2;" json:"rotation"`
}

var animation_frame_prop_positionTableInfo = &TableInfo{
	Name: "animation_frame_prop_position",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "animation_frame_id",
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
			GoFieldName:        "AnimationFrameID",
			GoFieldType:        "int32",
			JSONFieldName:      "animation_frame_id",
			ProtobufFieldName:  "animation_frame_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "part_type",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "USER_DEFINED",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "USER_DEFINED",
			ColumnLength:       -1,
			GoFieldName:        "PartType",
			GoFieldType:        "string",
			JSONFieldName:      "part_type",
			ProtobufFieldName:  "part_type",
			ProtobufType:       "string",
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

		&ColumnInfo{
			Index:              4,
			Name:               "rotation",
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
			GoFieldName:        "Rotation",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "rotation",
			ProtobufFieldName:  "rotation",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AnimationFramePropPosition) TableName() string {
	return "animation_frame_prop_position"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AnimationFramePropPosition) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AnimationFramePropPosition) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AnimationFramePropPosition) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AnimationFramePropPosition) TableInfo() *TableInfo {
	return animation_frame_prop_positionTableInfo
}
