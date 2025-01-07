// Code generated by ent, DO NOT EDIT.

package operation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the operation type in the database.
	Label = "operation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldDetail holds the string denoting the detail field in the database.
	FieldDetail = "detail"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldIsDone holds the string denoting the is_done field in the database.
	FieldIsDone = "is_done"
	// EdgeLabels holds the string denoting the labels edge name in mutations.
	EdgeLabels = "labels"
	// Table holds the table name of the operation in the database.
	Table = "operations"
	// LabelsTable is the table that holds the labels relation/edge. The primary key declared below.
	LabelsTable = "operation_labels"
	// LabelsInverseTable is the table name for the Label entity.
	// It exists in this package in order to avoid circular dependency with the "label" package.
	LabelsInverseTable = "labels"
)

// Columns holds all SQL columns for operation fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldDetail,
	FieldStatus,
	FieldIsDone,
}

var (
	// LabelsPrimaryKey and LabelsColumn2 are the table columns denoting the
	// primary key for the labels relation (M2M).
	LabelsPrimaryKey = []string{"operation_id", "label_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)

// OrderOption defines the ordering options for the Operation queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByIsDone orders the results by the is_done field.
func ByIsDone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDone, opts...).ToFunc()
}

// ByLabelsCount orders the results by labels count.
func ByLabelsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLabelsStep(), opts...)
	}
}

// ByLabels orders the results by labels terms.
func ByLabels(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLabelsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newLabelsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LabelsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, LabelsTable, LabelsPrimaryKey...),
	)
}
