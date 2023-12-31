// Code generated by ent, DO NOT EDIT.

package patient

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the patient type in the database.
	Label = "patient"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldFatherName holds the string denoting the father_name field in the database.
	FieldFatherName = "father_name"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldNationalCode holds the string denoting the national_code field in the database.
	FieldNationalCode = "national_code"
	// FieldIdentityCode holds the string denoting the identity_code field in the database.
	FieldIdentityCode = "identity_code"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeAttachment holds the string denoting the attachment edge name in mutations.
	EdgeAttachment = "attachment"
	// EdgeReception holds the string denoting the reception edge name in mutations.
	EdgeReception = "reception"
	// Table holds the table name of the patient in the database.
	Table = "patients"
	// AttachmentTable is the table that holds the attachment relation/edge.
	AttachmentTable = "attachments"
	// AttachmentInverseTable is the table name for the Attachment entity.
	// It exists in this package in order to avoid circular dependency with the "attachment" package.
	AttachmentInverseTable = "attachments"
	// AttachmentColumn is the table column denoting the attachment relation/edge.
	AttachmentColumn = "patient_attachment"
	// ReceptionTable is the table that holds the reception relation/edge.
	ReceptionTable = "receptions"
	// ReceptionInverseTable is the table name for the Reception entity.
	// It exists in this package in order to avoid circular dependency with the "reception" package.
	ReceptionInverseTable = "receptions"
	// ReceptionColumn is the table column denoting the reception relation/edge.
	ReceptionColumn = "patient_reception"
)

// Columns holds all SQL columns for patient fields.
var Columns = []string{
	FieldID,
	FieldFirstName,
	FieldLastName,
	FieldFatherName,
	FieldPhone,
	FieldNationalCode,
	FieldIdentityCode,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
)

// OrderOption defines the ordering options for the Patient queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFirstName orders the results by the first_name field.
func ByFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstName, opts...).ToFunc()
}

// ByLastName orders the results by the last_name field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByFatherName orders the results by the father_name field.
func ByFatherName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFatherName, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByNationalCode orders the results by the national_code field.
func ByNationalCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNationalCode, opts...).ToFunc()
}

// ByIdentityCode orders the results by the identity_code field.
func ByIdentityCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIdentityCode, opts...).ToFunc()
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

// ByAttachmentCount orders the results by attachment count.
func ByAttachmentCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAttachmentStep(), opts...)
	}
}

// ByAttachment orders the results by attachment terms.
func ByAttachment(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAttachmentStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByReceptionCount orders the results by reception count.
func ByReceptionCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newReceptionStep(), opts...)
	}
}

// ByReception orders the results by reception terms.
func ByReception(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReceptionStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAttachmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AttachmentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AttachmentTable, AttachmentColumn),
	)
}
func newReceptionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReceptionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ReceptionTable, ReceptionColumn),
	)
}
