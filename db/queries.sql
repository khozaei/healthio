-- name: GetPatient :one
SELECT * FROM patients 
WHERE id = ? LIMIT 1;

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

-- name: UpdateUser :exec
UPDATE users SET
    user_name = ?,
    password = ?
    WHERE id = ?;
