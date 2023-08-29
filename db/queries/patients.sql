-- name: GetPatient :one
SELECT * FROM patients 
WHERE id = ? LIMIT 1;

-- name: FindPatient :many
SELECT * FROM patients
WHERE deleted_at IS NULL AND 
    (last_name LIKE CONCAT('%',?,'%') OR
    first_name LIKE CONCAT('%',?,'%') OR
    national_code LIKE CONCAT('%',?,'%'))
ORDER BY last_name LIMIT ? OFFSET ?;

-- name: UpdatePatient :exec
UPDATE patients SET 
    first_name = ?,
    last_name = ?,
    national_code = ?,
    identity_code = ?,
    father_name = ?,
    phone = ?,
    updated_at = ?
    WHERE id = ?; 

-- name: AddPatient :execresult
INSERT INTO patients(first_name, last_name, father_name, national_code, phone,
            identity_code, created_at)
VALUES(?, ?, ?, ?, ?, ?, ?);

-- name: DeletePatient :exec
UPDATE patients SET
    deleted_at = ?
WHERE id = ?;

