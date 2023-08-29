// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: visits.sql

package model

import (
	"context"
	"database/sql"
	"time"
)

const addVisit = `-- name: AddVisit :execresult
INSERT INTO visits(visit_date, visit_time, visit_price,
            payment_type, is_paid, created_at, reception_id)
VALUES(?, ?, ?, ?, ?, ?, ?)
`

type AddVisitParams struct {
	VisitDate   sql.NullTime
	VisitTime   sql.NullTime
	VisitPrice  sql.NullInt64
	PaymentType sql.NullString
	IsPaid      sql.NullBool
	CreatedAt   time.Time
	ReceptionID int64
}

func (q *Queries) AddVisit(ctx context.Context, arg AddVisitParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addVisit,
		arg.VisitDate,
		arg.VisitTime,
		arg.VisitPrice,
		arg.PaymentType,
		arg.IsPaid,
		arg.CreatedAt,
		arg.ReceptionID,
	)
}

const deleteVisit = `-- name: DeleteVisit :exec
UPDATE visits SET
    deleted_at = ?
WHERE id = ?
`

type DeleteVisitParams struct {
	DeletedAt sql.NullTime
	ID        int64
}

func (q *Queries) DeleteVisit(ctx context.Context, arg DeleteVisitParams) error {
	_, err := q.db.ExecContext(ctx, deleteVisit, arg.DeletedAt, arg.ID)
	return err
}

const findVisits = `-- name: FindVisits :many
SELECT v.id, visit_price, visit_date, visit_time, payment_type, is_paid, reception_id, v.created_at, v.updated_at, v.deleted_at, r.id, reception_date, reception_time, visit_duration, insurance_code, description, r.created_at, r.updated_at, r.deleted_at, patient_id FROM visits AS v JOIN receptions AS r on r.id = v.reception_id 
WHERE v.deleted_at IS NULL AND
    v.visit_date BETWEEN ? AND ? 
ORDER BY v.visit_date, v.visit_time LIMIT ? OFFSET ?
`

type FindVisitsParams struct {
	FromVisitDate sql.NullTime
	ToVisitDate   sql.NullTime
	Limit         int32
	Offset        int32
}

type FindVisitsRow struct {
	ID            int64
	VisitPrice    sql.NullInt64
	VisitDate     sql.NullTime
	VisitTime     sql.NullTime
	PaymentType   sql.NullString
	IsPaid        sql.NullBool
	ReceptionID   int64
	CreatedAt     time.Time
	UpdatedAt     sql.NullTime
	DeletedAt     sql.NullTime
	ID_2          int64
	ReceptionDate time.Time
	ReceptionTime time.Time
	VisitDuration int32
	InsuranceCode sql.NullString
	Description   sql.NullString
	CreatedAt_2   time.Time
	UpdatedAt_2   sql.NullTime
	DeletedAt_2   sql.NullTime
	PatientID     int64
}

func (q *Queries) FindVisits(ctx context.Context, arg FindVisitsParams) ([]FindVisitsRow, error) {
	rows, err := q.db.QueryContext(ctx, findVisits,
		arg.FromVisitDate,
		arg.ToVisitDate,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindVisitsRow
	for rows.Next() {
		var i FindVisitsRow
		if err := rows.Scan(
			&i.ID,
			&i.VisitPrice,
			&i.VisitDate,
			&i.VisitTime,
			&i.PaymentType,
			&i.IsPaid,
			&i.ReceptionID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.ID_2,
			&i.ReceptionDate,
			&i.ReceptionTime,
			&i.VisitDuration,
			&i.InsuranceCode,
			&i.Description,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.DeletedAt_2,
			&i.PatientID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findVisitsByPatient = `-- name: FindVisitsByPatient :many
SELECT v.id, visit_price, visit_date, visit_time, payment_type, is_paid, reception_id, v.created_at, v.updated_at, v.deleted_at, r.id, reception_date, reception_time, visit_duration, insurance_code, description, r.created_at, r.updated_at, r.deleted_at, patient_id, p.id, first_name, last_name, father_name, phone, national_code, identity_code, p.created_at, p.updated_at, p.deleted_at FROM visits AS v JOIN receptions AS r on r.id = v.reception_id
    JOIN patients AS p on r.patient_id = p.id
WHERE p.deleted_at IS NULL AND
    v.deleted_at IS NULL AND 
    r.deleted_at IS NULL AND 
    (p.last_name LIKE CONCAT('%',?,'%') OR
    p.first_name LIKE CONCAT('%',?,'%') OR
    p.national_code LIKE CONCAT('%',?,'%')) AND
    v.visit_date BETWEEN ? AND ? 
ORDER BY v.visit_date, v.visit_time LIMIT ? OFFSET ?
`

type FindVisitsByPatientParams struct {
	CONCAT        interface{}
	CONCAT_2      interface{}
	CONCAT_3      interface{}
	FromVisitDate sql.NullTime
	ToVisitDate   sql.NullTime
	Limit         int32
	Offset        int32
}

type FindVisitsByPatientRow struct {
	ID            int64
	VisitPrice    sql.NullInt64
	VisitDate     sql.NullTime
	VisitTime     sql.NullTime
	PaymentType   sql.NullString
	IsPaid        sql.NullBool
	ReceptionID   int64
	CreatedAt     time.Time
	UpdatedAt     sql.NullTime
	DeletedAt     sql.NullTime
	ID_2          int64
	ReceptionDate time.Time
	ReceptionTime time.Time
	VisitDuration int32
	InsuranceCode sql.NullString
	Description   sql.NullString
	CreatedAt_2   time.Time
	UpdatedAt_2   sql.NullTime
	DeletedAt_2   sql.NullTime
	PatientID     int64
	ID_3          int64
	FirstName     sql.NullString
	LastName      sql.NullString
	FatherName    sql.NullString
	Phone         sql.NullString
	NationalCode  sql.NullString
	IdentityCode  sql.NullString
	CreatedAt_3   time.Time
	UpdatedAt_3   sql.NullTime
	DeletedAt_3   sql.NullTime
}

func (q *Queries) FindVisitsByPatient(ctx context.Context, arg FindVisitsByPatientParams) ([]FindVisitsByPatientRow, error) {
	rows, err := q.db.QueryContext(ctx, findVisitsByPatient,
		arg.CONCAT,
		arg.CONCAT_2,
		arg.CONCAT_3,
		arg.FromVisitDate,
		arg.ToVisitDate,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindVisitsByPatientRow
	for rows.Next() {
		var i FindVisitsByPatientRow
		if err := rows.Scan(
			&i.ID,
			&i.VisitPrice,
			&i.VisitDate,
			&i.VisitTime,
			&i.PaymentType,
			&i.IsPaid,
			&i.ReceptionID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.ID_2,
			&i.ReceptionDate,
			&i.ReceptionTime,
			&i.VisitDuration,
			&i.InsuranceCode,
			&i.Description,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.DeletedAt_2,
			&i.PatientID,
			&i.ID_3,
			&i.FirstName,
			&i.LastName,
			&i.FatherName,
			&i.Phone,
			&i.NationalCode,
			&i.IdentityCode,
			&i.CreatedAt_3,
			&i.UpdatedAt_3,
			&i.DeletedAt_3,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getVist = `-- name: GetVist :one
SELECT v.id, visit_price, visit_date, visit_time, payment_type, is_paid, reception_id, v.created_at, v.updated_at, v.deleted_at, r.id, reception_date, reception_time, visit_duration, insurance_code, description, r.created_at, r.updated_at, r.deleted_at, patient_id FROM visits AS v JOIN receptions AS r on r.id = v.reception_id 
WHERE v.id = ?
`

type GetVistRow struct {
	ID            int64
	VisitPrice    sql.NullInt64
	VisitDate     sql.NullTime
	VisitTime     sql.NullTime
	PaymentType   sql.NullString
	IsPaid        sql.NullBool
	ReceptionID   int64
	CreatedAt     time.Time
	UpdatedAt     sql.NullTime
	DeletedAt     sql.NullTime
	ID_2          int64
	ReceptionDate time.Time
	ReceptionTime time.Time
	VisitDuration int32
	InsuranceCode sql.NullString
	Description   sql.NullString
	CreatedAt_2   time.Time
	UpdatedAt_2   sql.NullTime
	DeletedAt_2   sql.NullTime
	PatientID     int64
}

func (q *Queries) GetVist(ctx context.Context, id int64) (GetVistRow, error) {
	row := q.db.QueryRowContext(ctx, getVist, id)
	var i GetVistRow
	err := row.Scan(
		&i.ID,
		&i.VisitPrice,
		&i.VisitDate,
		&i.VisitTime,
		&i.PaymentType,
		&i.IsPaid,
		&i.ReceptionID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.ID_2,
		&i.ReceptionDate,
		&i.ReceptionTime,
		&i.VisitDuration,
		&i.InsuranceCode,
		&i.Description,
		&i.CreatedAt_2,
		&i.UpdatedAt_2,
		&i.DeletedAt_2,
		&i.PatientID,
	)
	return i, err
}

const updateVisit = `-- name: UpdateVisit :exec
UPDATE visits SET
    visit_date = ?,
    visit_time = ?,
    visit_price = ?,
    payment_type = ?,
    is_paid = ?,
    updated_at = ?
WHERE id = ?
`

type UpdateVisitParams struct {
	VisitDate   sql.NullTime
	VisitTime   sql.NullTime
	VisitPrice  sql.NullInt64
	PaymentType sql.NullString
	IsPaid      sql.NullBool
	UpdatedAt   sql.NullTime
	ID          int64
}

func (q *Queries) UpdateVisit(ctx context.Context, arg UpdateVisitParams) error {
	_, err := q.db.ExecContext(ctx, updateVisit,
		arg.VisitDate,
		arg.VisitTime,
		arg.VisitPrice,
		arg.PaymentType,
		arg.IsPaid,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}