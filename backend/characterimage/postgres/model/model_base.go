package model

import "fmt"

// Action CRUD actions
type Action int32

var (
	// Create action when record is created
	Create = Action(0)

	// RetrieveOne action when a record is retrieved from db
	RetrieveOne = Action(1)

	// RetrieveMany action when record(s) are retrieved from db
	RetrieveMany = Action(2)

	// Update action when record is updated in db
	Update = Action(3)

	// Delete action when record is deleted in db
	Delete = Action(4)

	// FetchDDL action when fetching ddl info from db
	FetchDDL = Action(5)

	tables map[string]*TableInfo
)

func init() {
	tables = make(map[string]*TableInfo)

	tables["animation"] = animationTableInfo
	tables["animation_frame"] = animation_frameTableInfo
	tables["animation_frame_pixel"] = animation_frame_pixelTableInfo
	tables["animation_frame_prop_position"] = animation_frame_prop_positionTableInfo
	tables["animation_frame_static_position"] = animation_frame_static_positionTableInfo
	tables["body_type"] = body_typeTableInfo
	tables["color_string"] = color_stringTableInfo
	tables["dynamic_part"] = dynamic_partTableInfo
	tables["dynamic_part_mapping"] = dynamic_part_mappingTableInfo
	tables["dynamic_part_mapping_pixel"] = dynamic_part_mapping_pixelTableInfo
	tables["dynamic_part_pixel"] = dynamic_part_pixelTableInfo
	tables["prop"] = propTableInfo
	tables["prop_image"] = prop_imageTableInfo
	tables["schema_migrations"] = schema_migrationsTableInfo
	tables["static_part"] = static_partTableInfo
	tables["static_part_image"] = static_part_imageTableInfo
}

// String describe the action
func (i Action) String() string {
	switch i {
	case Create:
		return "Create"
	case RetrieveOne:
		return "RetrieveOne"
	case RetrieveMany:
		return "RetrieveMany"
	case Update:
		return "Update"
	case Delete:
		return "Delete"
	case FetchDDL:
		return "FetchDDL"
	default:
		return fmt.Sprintf("unknown action: %d", int(i))
	}
}

// Model interface methods for database structs generated
type Model interface {
	TableName() string
	BeforeSave() error
	Prepare()
	Validate(action Action) error
	TableInfo() *TableInfo
}

// TableInfo describes a table in the database
type TableInfo struct {
	Name    string        `json:"name"`
	Columns []*ColumnInfo `json:"columns"`
}

// ColumnInfo describes a column in the database table
type ColumnInfo struct {
	Index              int    `json:"index"`
	GoFieldName        string `json:"go_field_name"`
	GoFieldType        string `json:"go_field_type"`
	JSONFieldName      string `json:"json_field_name"`
	ProtobufFieldName  string `json:"protobuf_field_name"`
	ProtobufType       string `json:"protobuf_field_type"`
	ProtobufPos        int    `json:"protobuf_field_pos"`
	Comment            string `json:"comment"`
	Notes              string `json:"notes"`
	Name               string `json:"name"`
	Nullable           bool   `json:"is_nullable"`
	DatabaseTypeName   string `json:"database_type_name"`
	DatabaseTypePretty string `json:"database_type_pretty"`
	IsPrimaryKey       bool   `json:"is_primary_key"`
	IsAutoIncrement    bool   `json:"is_auto_increment"`
	IsArray            bool   `json:"is_array"`
	ColumnType         string `json:"column_type"`
	ColumnLength       int64  `json:"column_length"`
	DefaultValue       string `json:"default_value"`
}

// GetTableInfo retrieve TableInfo for a table
func GetTableInfo(name string) (*TableInfo, bool) {
	val, ok := tables[name]
	return val, ok
}
