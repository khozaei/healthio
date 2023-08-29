-- name: GetAttachment :one
SELECT * FROM attachments AS a JOIN visits AS v on v.id = a.visit_id 
WHERE a.id = ?;

-- name: FindAttachments :many
SELECT * FROM attachments AS a JOIN visits AS v on v.id = a.visit_id 
WHERE a.deleted_at IS NULL AND
    v.id = ? 
ORDER BY v.visit_date, v.visit_time LIMIT ? OFFSET ?;

-- name: FindAttachmentsByPatient :many
SELECT * FROM attachments AS a JOIN patients AS p on p.id = a.patient_id 
WHERE a.deleted_at IS NULL AND
    p.id = ? 
ORDER BY a.created_at LIMIT ? OFFSET ?;

-- name: UpdateAttachment :exec
UPDATE attachments SET
    file_path = ?,
    file_type = ?,
    updated_at = ?
WHERE id = ?;

-- name: DeleteAttachment :exec
UPDATE attachments SET
    deleted_at = ?
WHERE id = ?;

-- name: AddAttachment :execresult
INSERT INTO attachments(file_path, file_type, visit_id,
            patient_id, created_at)
VALUES(?, ?, ?, ?, ?);
