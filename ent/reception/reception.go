// Code generated by ent, DO NOT EDIT.

package reception

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the reception type in the database.
	Label = "reception"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldReceptionFor holds the string denoting the reception_for field in the database.
	FieldReceptionFor = "reception_for"
	// FieldVisitDuration holds the string denoting the visit_duration field in the database.
	FieldVisitDuration = "visit_duration"
	// FieldInsuranceCode holds the string denoting the insurance_code field in the database.
	FieldInsuranceCode = "insurance_code"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "patient"
	// EdgeVisit holds the string denoting the visit edge name in mutations.
	EdgeVisit = "visit"
	// Table holds the table name of the reception in the database.
	Table = "receptions"
	// PatientTable is the table that holds the patient relation/edge.
	PatientTable = "receptions"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the patient relation/edge.
	PatientColumn = "patient_reception"
	// VisitTable is the table that holds the visit relation/edge.
	VisitTable = "visits"
	// VisitInverseTable is the table name for the Visit entity.
	// It exists in this package in order to avoid circular dependency with the "visit" package.
	VisitInverseTable = "visits"
	// VisitColumn is the table column denoting the visit relation/edge.
	VisitColumn = "reception_visit"
)

// Columns holds all SQL columns for reception fields.
var Columns = []string{
	FieldID,
	FieldReceptionFor,
	FieldVisitDuration,
	FieldInsuranceCode,
	FieldDescription,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "receptions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"patient_reception",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// VisitDurationValidator is a validator for the "visit_duration" field. It is called by the builders before save.
	VisitDurationValidator func(int) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
)

// OrderOption defines the ordering options for the Reception queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByReceptionFor orders the results by the reception_for field.
func ByReceptionFor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReceptionFor, opts...).ToFunc()
}

// ByVisitDuration orders the results by the visit_duration field.
func ByVisitDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVisitDuration, opts...).ToFunc()
}

// ByInsuranceCode orders the results by the insurance_code field.
func ByInsuranceCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInsuranceCode, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByPatientField orders the results by patient field.
func ByPatientField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPatientStep(), sql.OrderByField(field, opts...))
	}
}

// ByVisitCount orders the results by visit count.
func ByVisitCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVisitStep(), opts...)
	}
}

// ByVisit orders the results by visit terms.
func ByVisit(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVisitStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPatientStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PatientInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
	)
}
func newVisitStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VisitInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, VisitTable, VisitColumn),
	)
}
