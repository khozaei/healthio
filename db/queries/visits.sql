-- name: GetVist :one
SELECT * FROM visits AS v JOIN receptions AS r on r.id = v.reception_id 
WHERE v.id = ?;

-- name: FindVisits :many
SELECT * FROM visits AS v JOIN receptions AS r on r.id = v.reception_id 
WHERE v.deleted_at IS NULL AND
    v.visit_date BETWEEN ? AND ? 
ORDER BY v.visit_date, v.visit_time LIMIT ? OFFSET ?;

-- name: FindVisitsByPatient :many
SELECT * FROM visits AS v JOIN receptions AS r on r.id = v.reception_id
    JOIN patients AS p on r.patient_id = p.id
WHERE p.deleted_at IS NULL AND
    v.deleted_at IS NULL AND 
    r.deleted_at IS NULL AND 
    (p.last_name LIKE CONCAT('%',?,'%') OR
    p.first_name LIKE CONCAT('%',?,'%') OR
    p.national_code LIKE CONCAT('%',?,'%')) AND
    v.visit_date BETWEEN ? AND ? 
ORDER BY v.visit_date, v.visit_time LIMIT ? OFFSET ?;

-- name: UpdateVisit :exec
UPDATE visits SET
    visit_date = ?,
    visit_time = ?,
    visit_price = ?,
    payment_type = ?,
    is_paid = ?,
    updated_at = ?
WHERE id = ?;

-- name: DeleteVisit :exec
UPDATE visits SET
    deleted_at = ?
WHERE id = ?;

-- name: AddVisit :execresult
INSERT INTO visits(visit_date, visit_time, visit_price,
            payment_type, is_paid, created_at, reception_id)
VALUES(?, ?, ?, ?, ?, ?, ?);
