-- name: insert-task
INSERT INTO maintenance.tasks (user_id, summary)
VALUES (?, AES_ENCRYPT(?, 'secure_key'));

-- name: update-task
UPDATE maintenance.tasks SET
    summary = AES_ENCRYPT(?, 'secure_key'),
    updated_at = NOW()
WHERE id = ?;

-- name: delete-task
DELETE FROM maintenance.tasks WHERE id = ?;

-- name: list-tasks
SELECT
    id AS ID,
    user_id AS UserID,
    AES_DECRYPT(summary, 'secure_key') AS Summary,
    created_at AS CreatedAt,
    updated_at AS UpdatedAt
FROM maintenance.tasks WHERE user_id = ?;

-- name: get-task
SELECT
    id AS ID,
    user_id AS UserID,
    AES_DECRYPT(summary, 'secure_key') AS Summary,
    created_at AS CreatedAt,
    updated_at AS UpdatedAt
FROM maintenance.tasks WHERE id = ?;

