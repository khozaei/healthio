// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khozaei/healthio/ent/attachment"
	"github.com/khozaei/healthio/ent/patient"
	"github.com/khozaei/healthio/ent/predicate"
	"github.com/khozaei/healthio/ent/visit"
)

// AttachmentUpdate is the builder for updating Attachment entities.
type AttachmentUpdate struct {
	config
	hooks    []Hook
	mutation *AttachmentMutation
}

// Where appends a list predicates to the AttachmentUpdate builder.
func (au *AttachmentUpdate) Where(ps ...predicate.Attachment) *AttachmentUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetFilePath sets the "file_path" field.
func (au *AttachmentUpdate) SetFilePath(s string) *AttachmentUpdate {
	au.mutation.SetFilePath(s)
	return au
}

// SetNillableFilePath sets the "file_path" field if the given value is not nil.
func (au *AttachmentUpdate) SetNillableFilePath(s *string) *AttachmentUpdate {
	if s != nil {
		au.SetFilePath(*s)
	}
	return au
}

// ClearFilePath clears the value of the "file_path" field.
func (au *AttachmentUpdate) ClearFilePath() *AttachmentUpdate {
	au.mutation.ClearFilePath()
	return au
}

// SetFileType sets the "file_type" field.
func (au *AttachmentUpdate) SetFileType(s string) *AttachmentUpdate {
	au.mutation.SetFileType(s)
	return au
}

// SetNillableFileType sets the "file_type" field if the given value is not nil.
func (au *AttachmentUpdate) SetNillableFileType(s *string) *AttachmentUpdate {
	if s != nil {
		au.SetFileType(*s)
	}
	return au
}

// ClearFileType clears the value of the "file_type" field.
func (au *AttachmentUpdate) ClearFileType() *AttachmentUpdate {
	au.mutation.ClearFileType()
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AttachmentUpdate) SetCreatedAt(t time.Time) *AttachmentUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AttachmentUpdate) SetNillableCreatedAt(t *time.Time) *AttachmentUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AttachmentUpdate) SetUpdatedAt(t time.Time) *AttachmentUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (au *AttachmentUpdate) SetNillableUpdatedAt(t *time.Time) *AttachmentUpdate {
	if t != nil {
		au.SetUpdatedAt(*t)
	}
	return au
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (au *AttachmentUpdate) ClearUpdatedAt() *AttachmentUpdate {
	au.mutation.ClearUpdatedAt()
	return au
}

// SetDeletedAt sets the "deleted_at" field.
func (au *AttachmentUpdate) SetDeletedAt(t time.Time) *AttachmentUpdate {
	au.mutation.SetDeletedAt(t)
	return au
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (au *AttachmentUpdate) SetNillableDeletedAt(t *time.Time) *AttachmentUpdate {
	if t != nil {
		au.SetDeletedAt(*t)
	}
	return au
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (au *AttachmentUpdate) ClearDeletedAt() *AttachmentUpdate {
	au.mutation.ClearDeletedAt()
	return au
}

// SetVisitID sets the "visit" edge to the Visit entity by ID.
func (au *AttachmentUpdate) SetVisitID(id int) *AttachmentUpdate {
	au.mutation.SetVisitID(id)
	return au
}

// SetNillableVisitID sets the "visit" edge to the Visit entity by ID if the given value is not nil.
func (au *AttachmentUpdate) SetNillableVisitID(id *int) *AttachmentUpdate {
	if id != nil {
		au = au.SetVisitID(*id)
	}
	return au
}

// SetVisit sets the "visit" edge to the Visit entity.
func (au *AttachmentUpdate) SetVisit(v *Visit) *AttachmentUpdate {
	return au.SetVisitID(v.ID)
}

// SetPatientID sets the "patient" edge to the Patient entity by ID.
func (au *AttachmentUpdate) SetPatientID(id int) *AttachmentUpdate {
	au.mutation.SetPatientID(id)
	return au
}

// SetNillablePatientID sets the "patient" edge to the Patient entity by ID if the given value is not nil.
func (au *AttachmentUpdate) SetNillablePatientID(id *int) *AttachmentUpdate {
	if id != nil {
		au = au.SetPatientID(*id)
	}
	return au
}

// SetPatient sets the "patient" edge to the Patient entity.
func (au *AttachmentUpdate) SetPatient(p *Patient) *AttachmentUpdate {
	return au.SetPatientID(p.ID)
}

// Mutation returns the AttachmentMutation object of the builder.
func (au *AttachmentUpdate) Mutation() *AttachmentMutation {
	return au.mutation
}

// ClearVisit clears the "visit" edge to the Visit entity.
func (au *AttachmentUpdate) ClearVisit() *AttachmentUpdate {
	au.mutation.ClearVisit()
	return au
}

// ClearPatient clears the "patient" edge to the Patient entity.
func (au *AttachmentUpdate) ClearPatient() *AttachmentUpdate {
	au.mutation.ClearPatient()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AttachmentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AttachmentUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AttachmentUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AttachmentUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AttachmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(attachment.Table, attachment.Columns, sqlgraph.NewFieldSpec(attachment.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.FilePath(); ok {
		_spec.SetField(attachment.FieldFilePath, field.TypeString, value)
	}
	if au.mutation.FilePathCleared() {
		_spec.ClearField(attachment.FieldFilePath, field.TypeString)
	}
	if value, ok := au.mutation.FileType(); ok {
		_spec.SetField(attachment.FieldFileType, field.TypeString, value)
	}
	if au.mutation.FileTypeCleared() {
		_spec.ClearField(attachment.FieldFileType, field.TypeString)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(attachment.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(attachment.FieldUpdatedAt, field.TypeTime, value)
	}
	if au.mutation.UpdatedAtCleared() {
		_spec.ClearField(attachment.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := au.mutation.DeletedAt(); ok {
		_spec.SetField(attachment.FieldDeletedAt, field.TypeTime, value)
	}
	if au.mutation.DeletedAtCleared() {
		_spec.ClearField(attachment.FieldDeletedAt, field.TypeTime)
	}
	if au.mutation.VisitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.VisitTable,
			Columns: []string{attachment.VisitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(visit.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.VisitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.VisitTable,
			Columns: []string{attachment.VisitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(visit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.PatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.PatientTable,
			Columns: []string{attachment.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(patient.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.PatientTable,
			Columns: []string{attachment.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(patient.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attachment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AttachmentUpdateOne is the builder for updating a single Attachment entity.
type AttachmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AttachmentMutation
}

// SetFilePath sets the "file_path" field.
func (auo *AttachmentUpdateOne) SetFilePath(s string) *AttachmentUpdateOne {
	auo.mutation.SetFilePath(s)
	return auo
}

// SetNillableFilePath sets the "file_path" field if the given value is not nil.
func (auo *AttachmentUpdateOne) SetNillableFilePath(s *string) *AttachmentUpdateOne {
	if s != nil {
		auo.SetFilePath(*s)
	}
	return auo
}

// ClearFilePath clears the value of the "file_path" field.
func (auo *AttachmentUpdateOne) ClearFilePath() *AttachmentUpdateOne {
	auo.mutation.ClearFilePath()
	return auo
}

// SetFileType sets the "file_type" field.
func (auo *AttachmentUpdateOne) SetFileType(s string) *AttachmentUpdateOne {
	auo.mutation.SetFileType(s)
	return auo
}

// SetNillableFileType sets the "file_type" field if the given value is not nil.
func (auo *AttachmentUpdateOne) SetNillableFileType(s *string) *AttachmentUpdateOne {
	if s != nil {
		auo.SetFileType(*s)
	}
	return auo
}

// ClearFileType clears the value of the "file_type" field.
func (auo *AttachmentUpdateOne) ClearFileType() *AttachmentUpdateOne {
	auo.mutation.ClearFileType()
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AttachmentUpdateOne) SetCreatedAt(t time.Time) *AttachmentUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AttachmentUpdateOne) SetNillableCreatedAt(t *time.Time) *AttachmentUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AttachmentUpdateOne) SetUpdatedAt(t time.Time) *AttachmentUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (auo *AttachmentUpdateOne) SetNillableUpdatedAt(t *time.Time) *AttachmentUpdateOne {
	if t != nil {
		auo.SetUpdatedAt(*t)
	}
	return auo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (auo *AttachmentUpdateOne) ClearUpdatedAt() *AttachmentUpdateOne {
	auo.mutation.ClearUpdatedAt()
	return auo
}

// SetDeletedAt sets the "deleted_at" field.
func (auo *AttachmentUpdateOne) SetDeletedAt(t time.Time) *AttachmentUpdateOne {
	auo.mutation.SetDeletedAt(t)
	return auo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auo *AttachmentUpdateOne) SetNillableDeletedAt(t *time.Time) *AttachmentUpdateOne {
	if t != nil {
		auo.SetDeletedAt(*t)
	}
	return auo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (auo *AttachmentUpdateOne) ClearDeletedAt() *AttachmentUpdateOne {
	auo.mutation.ClearDeletedAt()
	return auo
}

// SetVisitID sets the "visit" edge to the Visit entity by ID.
func (auo *AttachmentUpdateOne) SetVisitID(id int) *AttachmentUpdateOne {
	auo.mutation.SetVisitID(id)
	return auo
}

// SetNillableVisitID sets the "visit" edge to the Visit entity by ID if the given value is not nil.
func (auo *AttachmentUpdateOne) SetNillableVisitID(id *int) *AttachmentUpdateOne {
	if id != nil {
		auo = auo.SetVisitID(*id)
	}
	return auo
}

// SetVisit sets the "visit" edge to the Visit entity.
func (auo *AttachmentUpdateOne) SetVisit(v *Visit) *AttachmentUpdateOne {
	return auo.SetVisitID(v.ID)
}

// SetPatientID sets the "patient" edge to the Patient entity by ID.
func (auo *AttachmentUpdateOne) SetPatientID(id int) *AttachmentUpdateOne {
	auo.mutation.SetPatientID(id)
	return auo
}

// SetNillablePatientID sets the "patient" edge to the Patient entity by ID if the given value is not nil.
func (auo *AttachmentUpdateOne) SetNillablePatientID(id *int) *AttachmentUpdateOne {
	if id != nil {
		auo = auo.SetPatientID(*id)
	}
	return auo
}

// SetPatient sets the "patient" edge to the Patient entity.
func (auo *AttachmentUpdateOne) SetPatient(p *Patient) *AttachmentUpdateOne {
	return auo.SetPatientID(p.ID)
}

// Mutation returns the AttachmentMutation object of the builder.
func (auo *AttachmentUpdateOne) Mutation() *AttachmentMutation {
	return auo.mutation
}

// ClearVisit clears the "visit" edge to the Visit entity.
func (auo *AttachmentUpdateOne) ClearVisit() *AttachmentUpdateOne {
	auo.mutation.ClearVisit()
	return auo
}

// ClearPatient clears the "patient" edge to the Patient entity.
func (auo *AttachmentUpdateOne) ClearPatient() *AttachmentUpdateOne {
	auo.mutation.ClearPatient()
	return auo
}

// Where appends a list predicates to the AttachmentUpdate builder.
func (auo *AttachmentUpdateOne) Where(ps ...predicate.Attachment) *AttachmentUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AttachmentUpdateOne) Select(field string, fields ...string) *AttachmentUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Attachment entity.
func (auo *AttachmentUpdateOne) Save(ctx context.Context) (*Attachment, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AttachmentUpdateOne) SaveX(ctx context.Context) *Attachment {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AttachmentUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AttachmentUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AttachmentUpdateOne) sqlSave(ctx context.Context) (_node *Attachment, err error) {
	_spec := sqlgraph.NewUpdateSpec(attachment.Table, attachment.Columns, sqlgraph.NewFieldSpec(attachment.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Attachment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attachment.FieldID)
		for _, f := range fields {
			if !attachment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != attachment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.FilePath(); ok {
		_spec.SetField(attachment.FieldFilePath, field.TypeString, value)
	}
	if auo.mutation.FilePathCleared() {
		_spec.ClearField(attachment.FieldFilePath, field.TypeString)
	}
	if value, ok := auo.mutation.FileType(); ok {
		_spec.SetField(attachment.FieldFileType, field.TypeString, value)
	}
	if auo.mutation.FileTypeCleared() {
		_spec.ClearField(attachment.FieldFileType, field.TypeString)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(attachment.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(attachment.FieldUpdatedAt, field.TypeTime, value)
	}
	if auo.mutation.UpdatedAtCleared() {
		_spec.ClearField(attachment.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := auo.mutation.DeletedAt(); ok {
		_spec.SetField(attachment.FieldDeletedAt, field.TypeTime, value)
	}
	if auo.mutation.DeletedAtCleared() {
		_spec.ClearField(attachment.FieldDeletedAt, field.TypeTime)
	}
	if auo.mutation.VisitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.VisitTable,
			Columns: []string{attachment.VisitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(visit.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.VisitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.VisitTable,
			Columns: []string{attachment.VisitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(visit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.PatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.PatientTable,
			Columns: []string{attachment.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(patient.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.PatientTable,
			Columns: []string{attachment.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(patient.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Attachment{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attachment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}