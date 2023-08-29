-- name: GetReceptions :many
SELECT * FROM receptions AS r join patients AS p on r.patient_id = p.id 
WHERE r.deleted_at IS NULL AND
    r.reception_date BETWEEN ? AND ? 
ORDER BY r.reception_date LIMIT ? OFFSET ?;

-- name: FindReceptionsByPatient :many
SELECT * FROM receptions AS r join patients AS p on r.patient_id = p.id 
WHERE r.deleted_at IS NULL AND 
    p.deleted_at IS NULL AND 
    (p.last_name LIKE CONCAT('%',?,'%') OR
    p.first_name LIKE CONCAT('%',?,'%') OR
    p.national_code LIKE CONCAT('%',?,'%')) AND
    r.reception_date BETWEEN ? AND ? 
ORDER BY p.last_name LIMIT ? OFFSET ?;

-- name: UpdateReception :exec
UPDATE receptions SET
    reception_date = ?,
    reception_time = ?,
    visit_duration = ?,
    insurance_code = ?,
    description = ?,
    updated_at = ?
WHERE id = ?;

-- name: DeleteReception :exec
UPDATE receptions SET
    deleted_at = ?
WHERE id = ?;

-- name: AddReception :execresult
INSERT INTO receptions(reception_date, reception_time, visit_duration,
            insurance_code, description, created_at, patient_id)
VALUES(?, ?, ?, ?, ?, ?, ?);
