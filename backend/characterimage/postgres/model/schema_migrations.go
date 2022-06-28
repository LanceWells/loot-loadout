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


Table: schema_migrations
[ 0] version                                        INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
[ 1] dirty                                          BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []


JSON Sample
-------------------------------------
{    "version": 45,    "dirty": true}



*/

// SchemaMigrations struct is a row record of the schema_migrations table in the charimage database
type SchemaMigrations struct {
	//[ 0] version                                        INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
	Version int64 `gorm:"primary_key;column:version;type:INT8;" json:"version"`
	//[ 1] dirty                                          BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	Dirty bool `gorm:"column:dirty;type:BOOL;" json:"dirty"`
}

var schema_migrationsTableInfo = &TableInfo{
	Name: "schema_migrations",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "version",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "Version",
			GoFieldType:        "int64",
			JSONFieldName:      "version",
			ProtobufFieldName:  "version",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "dirty",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "Dirty",
			GoFieldType:        "bool",
			JSONFieldName:      "dirty",
			ProtobufFieldName:  "dirty",
			ProtobufType:       "bool",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SchemaMigrations) TableName() string {
	return "schema_migrations"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SchemaMigrations) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SchemaMigrations) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SchemaMigrations) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SchemaMigrations) TableInfo() *TableInfo {
	return schema_migrationsTableInfo
}
